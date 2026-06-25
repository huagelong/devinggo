import type {
  VbenFormSchema as FormSchema,
  VbenFormProps,
} from '@vben/common-ui';

import type { ComponentType } from './component';

import { setupVbenForm, useVbenForm as useForm, z } from '@vben/common-ui';
import { $t } from '@vben/locales';

async function initSetupVbenForm() {
  setupVbenForm<ComponentType>({
    config: {
      // tdesign组件库默认都是 v-model:value
      baseModelPropName: 'value',

      // 一些组件是 v-model:checked 或者 v-model:fileList
      modelPropNameMap: {
        Checkbox: 'checked',
        Radio: 'checked',
        Switch: 'checked',
        Upload: 'fileList',
      },
    },
    defineRules: {
      required: (value, _params, ctx) => {
        if (value === undefined || value === null || value === '') {
          return $t('ui.formRules.required', [ctx.label]);
        }
        if (Array.isArray(value) && value.length === 0) {
          return $t('ui.formRules.required', [ctx.label]);
        }
        return true;
      },
      selectRequired: (value, _params, ctx) => {
        if (value === undefined || value === null || value === '') {
          return $t('ui.formRules.selectRequired', [ctx.label]);
        }
        if (Array.isArray(value) && value.length === 0) {
          return $t('ui.formRules.selectRequired', [ctx.label]);
        }
        return true;
      },
      pattern: (value, params, ctx) => {
        if (value === undefined || value === null || value === '') {
          return true;
        }
        const rawParam = Array.isArray(params) ? params[0] : params?.[0] ?? params;
        if (typeof rawParam !== 'string') {
          return true;
        }
        const [patternStr = '', message] = rawParam.split('#');
        if (!patternStr) {
          return true;
        }
        const regex = new RegExp(patternStr);
        if (!regex.test(String(value))) {
          return message || $t('ui.formRules.formatInvalid', [ctx.label]);
        }
        return true;
      },
    },
  });
}

const useVbenForm = useForm<ComponentType>;

export { initSetupVbenForm, useVbenForm, z };

export type VbenFormSchema = FormSchema<ComponentType>;
export type { VbenFormProps };
