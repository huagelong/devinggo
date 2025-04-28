<template>
  <a-modal width="1200px" v-model:visible="visible" :footer="false">
    <template #title>管理配置</template>
    <div class="h-full">
      <ma-crud ref="crudRef" :options="crud" :columns="columns" />
    </div>
  </a-modal>
</template>

<script setup>
import { reactive, ref } from 'vue'
import config from '@/api/system/config'
import { inputComponent } from './js/configDefine.js'

const visible = ref(false)
const groupId = ref()
const crudRef = ref()

const open = (id) => {
  groupId.value = id
  visible.value = true
  crudRef.value.requestData()
}
const crud = reactive({
  pk: 'key',
  autoRequest: false,
  api: config.getConfigList,
  rowSelection: { showCheckedAll: true, key: 'key' },
  showIndex: false,
  operationColumn: true,
  operationColumnWidth: 160,
  edit: { show: true, api: config.update, auth: ['system:config:update'] },
  delete: { show: true, api: config.delete, auth: ['system:config:delete'] },
  beforeRequest: (params) => params.group_id = groupId.value,
  beforeEdit: (record) => {
    if (record.config_select_data && typeof record.config_select_data === 'string') {
      record.config_select_data = record.config_select_data.replace(/\r|\n|\s/g, '')
      record.config_select_data = record.config_select_data.replace(',]', ']')
    }
    return true
  },
  formOption: { width: '700px' }
})

const columns = reactive([
  {
    title: '所属组',
    dataIndex: 'group_id',
    formType: 'select',
    hide: true,
    dict: { url: 'system/setting/configGroup/index', props: { label: 'name', value: 'id' } },
    commonRules: [{ required: true, message: '所属组必选' }],
    width: 180,
  },
  {
    title: '配置标题',
    dataIndex: 'name',
    search: true,
    commonRules: [{ required: true, message: '配置标题必填' }],
    width: 220,
  },
  {
    title: '配置标识',
    dataIndex: 'key',
    search: true,
    commonRules: [{ required: true, message: '配置标识必填' }],
    disabled: true,
    width: 180,
  },
  {
    title: '配置值',
    dataIndex: 'value',
    width: 200,
  },
  {
    title: '排序',
    dataIndex: 'sort',
    formType: 'input-number',
    width: 200,
    min: 0,
    max: 999,
  },
  {
    title: '输入组件',
    dataIndex: 'input_type',
    formType: 'select',
    commonRules: [{ required: true, message: '输入组件必选' }],
    dict: { data: inputComponent },
    onControl: (val, maFormObject) => {
      const service = maFormObject.getColumnService()
      service.get('config_select_data').setAttr('display', ['select', 'radio', 'checkbox'].includes(val))
    },
    width: 180,
  },
  {
    title: '配置说明',
    dataIndex: 'remark',
    formType: 'input',
    width: 180,
  },
  {
    title: '配置选择数据',
    dataIndex: 'config_select_data',
    formType: 'code-editor',
    height: 200,
    hide: true,
    extra: '用于配置下拉、单选、复选的数据，格式例子：[{"label":"数据一", "value":"shuju1"},...]'
  }
])

defineExpose({ open })
</script>