<script lang="ts" setup>
import type { ConfigFieldMeta } from '../model';

import { computed } from 'vue';

import { $t } from '@vben/locales';

import {
  Checkbox,
  CheckboxGroup,
  Input,
  Radio,
  RadioGroup,
  Select,
  Switch,
  Textarea,
} from 'tdesign-vue-next';

import ConfigRichTextEditor from './config-rich-text-editor.vue';
import ConfigUploadInput from './config-upload-input.vue';
import KeyValueEditor from './key-value-editor.vue';

const props = defineProps<{
  field: ConfigFieldMeta;
  modelValue?: any;
}>();

const emit = defineEmits<{
  (e: 'update:modelValue', value: any): void;
}>();

const innerValue = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value),
});

const selectOptions = computed(
  () =>
    props.field.config_select_data?.map((item) => ({
      label: item.label,
      value: item.value,
    })) ?? [],
);

const switchValues = computed(() => {
  const defaults = props.field.switchValues ?? { checked: 1, unchecked: 0 };
  return [defaults.checked, defaults.unchecked];
});

const normalizedInputType = computed(() =>
  String(props.field.input_type ?? '')
    .trim()
    .toLowerCase(),
);
</script>

<template>
  <div class="config-field-renderer">
    <Input
      v-if="normalizedInputType === 'input'"
      v-model="innerValue"
      allow-clear
      :placeholder="$t('common.enterContent')"
    />
    <Textarea
      v-else-if="normalizedInputType === 'textarea'"
      v-model="innerValue"
      :autosize="{ minRows: 3, maxRows: 6 }"
      :placeholder="$t('common.enterContent')"
    />
    <ConfigRichTextEditor
      v-else-if="normalizedInputType === 'editor'"
      v-model="innerValue"
    />
    <Select
      v-else-if="normalizedInputType === 'select'"
      v-model="innerValue"
      :options="selectOptions"
      :placeholder="$t('ui.placeholder.select')"
      clearable
      class="w-full"
    />
    <RadioGroup
      v-else-if="normalizedInputType === 'radio'"
      v-model="innerValue"
      class="flex flex-wrap gap-3"
    >
      <Radio
        v-for="option in selectOptions"
        :key="option.value"
        :value="option.value"
      >
        {{ option.label }}
      </Radio>
    </RadioGroup>
    <CheckboxGroup
      v-else-if="normalizedInputType === 'checkbox'"
      v-model="innerValue"
      class="flex flex-wrap gap-3"
    >
      <Checkbox
        v-for="option in selectOptions"
        :key="option.value"
        :value="option.value"
      >
        {{ option.label }}
      </Checkbox>
    </CheckboxGroup>
    <Switch
      v-else-if="normalizedInputType === 'switch'"
      v-model="innerValue"
      :custom-value="switchValues"
    />
    <ConfigUploadInput
      v-else-if="normalizedInputType === 'upload'"
      v-model="innerValue"
    />
    <KeyValueEditor
      v-else-if="normalizedInputType === 'key-value'"
      v-model="innerValue"
    />
    <Input
      v-else
      v-model="innerValue"
      allow-clear
      :placeholder="$t('common.enterContent')"
    />
  </div>
</template>

<style scoped>
.config-field-renderer :deep(.t-input),
.config-field-renderer :deep(.t-textarea),
.config-field-renderer :deep(.t-select),
.config-field-renderer :deep(.t-tree-select),
.config-field-renderer :deep(.t-date-picker),
.config-field-renderer :deep(.t-range-input) {
  width: 100%;
}

.config-field-renderer :deep(.t-input),
.config-field-renderer :deep(.t-textarea__inner),
.config-field-renderer :deep(.t-select),
.config-field-renderer :deep(.t-select__wrap) {
  border-radius: 8px;
}
</style>
