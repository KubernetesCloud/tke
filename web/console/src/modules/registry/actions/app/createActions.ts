import {
  ReduxAction,
  extend,
  generateWorkflowActionCreator,
  OperationTrigger,
  isSuccessWorkflow,
  createFFObjectActions
} from '@tencent/ff-redux';
import { App, AppFilter, RootState, ChartInfo, ChartInfoFilter } from '../../models/index';
import * as ActionType from '../../constants/ActionType';
import * as WebAPI from '../../WebAPI';
import { initAppCreationState } from '../../constants/initState';
import { AppValidateSchema } from '../../constants/AppValidateConfig';
import { router } from '../../router';
import { createValidatorActions, getValidatorActionType } from '@tencent/ff-validator';
type GetState = () => RootState;
const routerSea = seajs.require('router');

/**
 * 增加仓库
 */
const addAppWorkflow = generateWorkflowActionCreator<App, AppFilter>({
  actionType: ActionType.AddApp,
  workflowStateLocator: (state: RootState) => state.appAddWorkflow,
  operationExecutor: WebAPI.addApp,
  after: {
    [OperationTrigger.Done]: (dispatch: Redux.Dispatch, getState: GetState) => {
      let { appAddWorkflow, route } = getState();
      if (isSuccessWorkflow(appAddWorkflow)) {
        let params = appAddWorkflow.params;
        /// #if project
        routerSea.navigate(`tkestack-project/app/app?cluster=${params.cluster}&namespace=${params.namespace}`);
        /// #else
        routerSea.navigate(
          `tkestack/application/app?cluster=${params.cluster}&namespace=${params.namespace}&projectId=${params.projectId}`
        );
        /// #endif
        //进入列表时自动加载
        //退出状态页面时自动清理状态
      }
      /** 结束工作流 */
      dispatch(createActions.addAppWorkflow.reset());
    }
  }
});

/**
 * 获取chart详情
 */
const fetchChartInfoActions = createFFObjectActions<ChartInfo, ChartInfoFilter>({
  actionName: ActionType.ChartInfo,
  fetcher: async query => {
    let response = await WebAPI.fetchChartInfo(query.filter);
    return response;
  },
  getRecord: (getState: GetState) => {
    return getState().chartInfo;
  },
  onFinish: (record, dispatch: Redux.Dispatch, getState: GetState) => {
    let { appCreation } = getState();
    let values = Object.assign({}, appCreation.spec.values);
    if (record.data && record.data.spec.values && record.data.spec.values['values.yaml']) {
      values.rawValues = record.data.spec.values['values.yaml'];
    } else {
      values.rawValues = '';
    }
    dispatch(createActions.updateCreationState({ spec: Object.assign({}, appCreation.spec, { values: values }) }));
  }
});

const restActions = {
  addAppWorkflow,

  validator: createValidatorActions({
    userDefinedSchema: AppValidateSchema,
    validateStateLocator: (store: RootState) => {
      return store.appCreation;
    },
    validatorStateLocation: (store: RootState) => {
      return store.appValidator;
    }
  }),

  /** 更新状态 */
  updateCreationState: obj => {
    return (dispatch: Redux.Dispatch, getState: GetState) => {
      let { appCreation } = getState();
      dispatch({
        type: ActionType.UpdateAppCreationState,
        payload: Object.assign({}, appCreation, obj)
      });
    };
  },

  /** 离开创建页面，清除Creation当中的内容 */
  clearCreationState: (): ReduxAction<any> => {
    return {
      type: ActionType.UpdateAppCreationState,
      payload: initAppCreationState
    };
  },

  /** 离开创建页面，清除Validator当中的内容 */
  clearValidatorState: (): ReduxAction<any> => {
    return {
      type: getValidatorActionType(AppValidateSchema.formKey),
      payload: {}
    };
  }
};

export const createActions = extend({}, { chart: fetchChartInfoActions }, restActions);
