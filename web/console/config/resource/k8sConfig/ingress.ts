import { DetailField, DisplayField, DetailInfo } from '../../../src/modules/common/models';
import {
  defaulNotExistedValue,
  commonActionField,
  dataFormatConfig,
  commonTabList,
  generateResourceInfo
} from '../common';
import { t, Trans } from '@tencent/tea-app/lib/i18n';

const displayField: DisplayField = {
  name: {
    dataField: ['metadata.name'],
    dataFormat: dataFormatConfig['text'],
    width: '10%',
    headTitle: t('名称'),
    noExsitedValue: defaulNotExistedValue,
    isLink: true, // 判断该值是否为链接,
    isClip: true
  },
  backendService: {
    dataField: ['spec.rules'],
    dataFormat: dataFormatConfig['ingressRule'],
    width: '20%',
    headTitle: t('后端服务'),
    noExsitedValue: defaulNotExistedValue
  },
  creationTimestamp: {
    dataField: ['metadata.creationTimestamp'],
    dataFormat: dataFormatConfig['time'],
    width: '20%',
    headTitle: t('创建时间'),
    noExsitedValue: defaulNotExistedValue
  },
  operator: {
    dataField: [''],
    dataFormat: dataFormatConfig['operator'],
    width: '15%',
    headTitle: t('操作'),
    operatorList: [
      {
        name: t('编辑YAML'),
        actionType: 'modify',
        isInMoreOp: false
      },
      {
        name: t('删除'),
        actionType: 'delete',
        isInMoreOp: false
      }
    ]
  }
};

/** resrouce action当中的配置 */
const actionField = Object.assign({}, commonActionField);

/** 自定义配置详情的展示 */
const detailBasicInfo: DetailInfo = {
  info: {
    metadata: {
      dataField: ['metadata'],
      displayField: {
        name: {
          dataField: ['name'],
          dataFormat: dataFormatConfig['text'],
          label: t('名称'),
          noExsitedValue: defaulNotExistedValue,
          order: '0'
        },
        namespace: {
          dataField: ['namespace'],
          dataFormat: dataFormatConfig['text'],
          label: 'Namespace',
          noExsitedValue: defaulNotExistedValue,
          order: '5'
        },
        description: {
          dataField: ['annotations.description'],
          dataFormat: dataFormatConfig['text'],
          label: t('描述'),
          noExsitedValue: defaulNotExistedValue,
          order: '6'
        },
        labels: {
          dataField: ['labels'],
          dataFormat: dataFormatConfig['labels'],
          label: 'Labels',
          noExsitedValue: defaulNotExistedValue,
          order: '10'
        },
        createTime: {
          dataField: ['creationTimestamp'],
          dataFormat: dataFormatConfig['time'],
          label: t('创建时间'),
          noExsitedValue: defaulNotExistedValue,
          order: '20'
        }
      }
    }
  },
  rules: {
    spec: {
      dataField: ['spec'],
      displayField: {
        rules: {
          dataField: ['rules'],
          dataFormat: dataFormatConfig['rules'],
          label: t('转发配置'),
          noExsitedValue: defaulNotExistedValue
        }
      }
    }
  }
};

/** 详情页面的相关配置 */
const detailField: DetailField = {
  tabList: [...commonTabList],
  detailInfo: Object.assign({}, detailBasicInfo)
};

/** ingress 的配置 */
export const ingress = (k8sVersion: string) => {
  return generateResourceInfo({
    k8sVersion,
    resourceName: 'ingress',
    isRelevantToNamespace: true,
    requestType: {
      list: 'ingresses'
    },
    displayField,
    actionField,
    detailField
  });
};
