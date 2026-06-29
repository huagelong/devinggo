<script lang="ts" setup>
import { $t } from '@vben/locales';

import { Button, Dialog, Form, FormItem, Select } from 'tdesign-vue-next';

import type { DictOption } from '#/composables/crud/use-dict-options';

interface Props {
  visible: boolean;
  selectedHomePage: string | undefined;
  homePageOptions: DictOption[];
  loading: boolean;
}

interface Emits {
  (e: 'update:visible', value: boolean): void;
  (e: 'update:selectedHomePage', value: string | undefined): void;
  (e: 'save'): void;
  (e: 'cancel'): void;
}

defineProps<Props>();
const emit = defineEmits<Emits>();

function handleUpdateVisible(value: boolean) {
  emit('update:visible', value);
}

function handleUpdateSelectedHomePage(value: any) {
  emit('update:selectedHomePage', value);
}

function handleSave() {
  emit('save');
}

function handleCancel() {
  emit('cancel');
}
</script>

<template>
  <Dialog
    :visible="visible"
    width="520px"
    :header="$t('system.user.setHomePage')"
    destroy-on-close
    @update:visible="handleUpdateVisible"
  >
    <Form label-width="130px" layout="inline">
      <FormItem :label="$t('system.user.selectHomePage')">
        <Select
          :value="selectedHomePage"
          :options="homePageOptions"
          :placeholder="$t('ui.placeholder.select', [$t('system.user.selectHomePage')])"
          clearable
          class="w-full"
          @change="handleUpdateSelectedHomePage"
        />
      </FormItem>
    </Form>
    <template #footer>
      <div class="flex justify-end gap-2">
        <Button theme="default" @click="handleCancel">{{ $t('common.cancel') }}</Button>
        <Button theme="primary" :loading="loading" @click="handleSave">
          {{ $t('common.save') }}
        </Button>
      </div>
    </template>
  </Dialog>
</template>