<template>
  <div class="ma-content-block lg:flex justify-between p-4">
    <!-- CRUD 组件 -->
    <ma-crud :options="crud" :columns="columns" ref="crudRef">
      <!-- 状态列 -->
      <template #status="{ record }">
        <a-switch :checked-value="1" unchecked-value="2" @change="changeStatus($event, record.id)"
          :default-checked="record.status == 1" />
      </template>
    </ma-crud>
  </div>
</template>

<script setup>
  import { ref, reactive, computed } from 'vue'
  import api from '@/api/system/api'
  import { Message } from '@arco-design/web-vue'

  const crudRef = ref()

  let isRecovery = computed(() => crudRef.value ? crudRef.value.isRecovery : false )

  const changeStatus = async (status, id) => {
    const response = await api.changeStatus({ id, status })
    if (response.success) {
      Message.success(response.message)
    }
  }

  const crud = reactive({
    api: api.getList,
    recycleApi: api.getRecycleList,
    showIndex: false,
    pageLayout: 'fixed',
    rowSelection: { showCheckedAll: true },
    operationColumn: true,
    operationColumnWidth: 260,
    add: { show: true, api: api.save, auth: ['system:api:save'] },
    edit: { show: true, api: api.update, auth: ['system:api:update'] },
    delete: {
      show: true,
      api: api.deletes, auth: ['system:api:delete'],
      realApi: api.realDeletes, realAuth: ['system:api:realDeletes']
    },
    recovery: { show: true, api: api.recoverys, auth: ['system:api:recovery']},
    formOption: {
      id: 'apiManage',
      width: '850px',
      viewType: 'tag',
      tagId: 'apiForm',
      tagName: '接口',
      titleDataIndex: 'name',
      layout: [
        {
          formType: 'tabs',
          tabs: [
            { 
              title: '基础信息',
              formList: [
                { dataIndex: 'group_id' },
                { dataIndex: 'name' },
                { dataIndex: 'access_name' },
                { dataIndex: 'request_mode' },
                { dataIndex: 'status' },
                { dataIndex: 'auth_mode' },
                { dataIndex: 'remark' },
              ]
            },
          ]  
        },
      ]
    }
  })

  const columns = reactive([
    {
      title: '所属组', dataIndex: 'group_id', search: true, commonRules: [{ required: true, message: '所属组必选' }],
      formType: 'select', dict: { url: 'system/apiGroup/list', props: { label: 'name', value: 'id' }, translation: true },
      width: 140
    },
    {
      title: '接口名称', dataIndex: 'name', search: true, commonRules: [{ required: true, message: '应用名称必填' }],
      width: 150,
    },
    {
      title: '接口标识', dataIndex: 'access_name', width: 140,
      commonRules: [{ required: true, message: '接口标识必填' }],
      extra: '例子：system:app:getAppSecret'
    },
    {
      title: '请求模式', dataIndex: 'request_mode', search: true, formType: 'select',
      commonRules: [{ required: true, message: '请求模式必选' }],
      dict: { name: 'request_mode', props: { label: 'title', value: 'key' }, translation: true },
      width: 140,
    },
    {
      title: '状态', dataIndex: 'status', search: true, formType: 'radio',
      dict: { name: 'data_status', props: { label: 'title', value: 'key' } },
      addDefaultValue: '1', width: 80,
    },
    {
      title: '认证模式', dataIndex: 'auth_mode', formType: 'radio',
      dict: { data: [{ label: '简易模式', value: 1 }, { label: '复杂模式', value: 2 }], translation: true },
      addDefaultValue: 1, width: 130
    },
    {
      title: '备注', dataIndex: 'remark', hide: true, formType: 'textarea',
    },
    {
      title: '创建时间', dataIndex: 'created_at', addDisplay: false, editDisplay: false,
      width: 180,
    },
  ])
</script>

<script>
export default { name: 'system:api' }
</script>

<style scoped></style>
