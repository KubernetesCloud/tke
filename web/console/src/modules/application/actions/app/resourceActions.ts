import { extend, createFFObjectActions, uuid } from '@tencent/ff-redux';
import * as JsYAML from 'js-yaml';
import { RootState, AppResource, AppResourceFilter, Resource } from '../../models';
import * as ActionTypes from '../../constants/ActionTypes';
import * as WebAPI from '../../WebAPI';
import { t, Trans } from '@tencent/tea-app/lib/i18n';
type GetState = () => RootState;
const tips = seajs.require('tips');

/**
 * 获取资源详情
 */

const fetchAppResourceActions = createFFObjectActions<AppResource, AppResourceFilter>({
  actionName: ActionTypes.AppResource,
  fetcher: async (query, getState: GetState, fetchOptions, dispatch: Redux.Dispatch) => {
    let response = await WebAPI.fetchAppResource(query.filter);
    return response;
  },
  getRecord: (getState: GetState) => {
    return getState().appResource;
  },
  onFinish: (record, dispatch: Redux.Dispatch, getState: GetState) => {
    let resources = [];
    if (record.data) {
      try {
        Object.keys(record.data.spec.resources).forEach(k => {
          record.data.spec.resources[k].forEach(item => {
            let json = JsYAML.safeLoad(item);
            resources.push({
              id: uuid(),
              metadata: {
                namespace: json.metadata.namespace,
                name: json.metadata.name
              },
              kind: json.kind,
              cluster: record.data.spec.targetCluster,
              yaml: JsYAML.safeDump(json)
            });
          });
        });
      } catch (e) {
        // console.log(e);
        tips.error(t('资源列表读取失败'), 2000);
      }
    }
    dispatch({
      type: ActionTypes.ResourceList,
      payload: {
        resources: resources
      }
    });
  }
});

const restActions = {};

export const resourceActions = extend({}, fetchAppResourceActions, restActions);
