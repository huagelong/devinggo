<!--vue.js template for generating resource-->
<!--@Link  https://github.com/huagelong/devinggo-->
<!--@Copyright  Copyright (c) 2024 devinggo-->
<!--@Author Kai <hpuwang@gmail.com>-->
<!--@License  https://github.com/huagelong/devinggo/blob/master/LICENSE-->

<template>
  <div class="ma-content-block lg:flex justify-between p-4">
    <!-- CRUD 组件 -->
    <ma-crud :options="options" :columns="columns" ref="crudRef">
      <template #status="{ record }">
        <div v-if="record.name == 'system'" >
          <a-switch disabled :checked-value="1" unchecked-value="2" @change="switchStatus($event, record)"
                    :default-checked="record.status == 1"/>
        </div>
        <div v-else>
          <a-switch :checked-value="1" unchecked-value="2" @change="switchStatus($event, record)"
                    :default-checked="record.status == 1"/>
        </div>
      </template>

      <template #operationCell="{ record }">
        <div v-if="record.name == 'system'"></div>
      </template>

    </ma-crud>
  </div>
</template>
<script setup>
import {ref, reactive} from 'vue'
import systemModules from '@/api/system/systemModules'
import {Message} from '@arco-design/web-vue'
import tool from '@/utils/tool'
import * as common from '@/utils/common'

const crudRef = ref()

const switchStatus = (status, record) => {
  systemModules.changeStatus({id: record.id, status}).then(res => {
    res.success && Message.success(res.message)
  }).catch(e => {
    console.log(e)
  })
}

const options = reactive({
  add: {
    api: systemModules.save,
    auth: [
      "system:systemModules:save"
    ],
    show: true
  },
  api: systemModules.getPageList,
  delete: {
    api: systemModules.deletes,
    auth: [
      "system:systemModules:delete"
    ],
    realApi: systemModules.realDeletes,
    realAuth: [
      "system:systemModules:realDelete"
    ],
    show: true
  },
  edit: {
    api: systemModules.update,
    auth: [
      "system:systemModules:update"
    ],
    show: true
  },
  formOption: {
    viewType: 'modal',
    width: 600
  },
  id: 'system_modules',
  operationColumn: true,
  operationColumnWidth: 160,
  pk: 'id',
  recovery: {
    api: systemModules.recoverys,
    auth: [
      "system:systemModules:recovery"
    ],
    show: true
  },
  recycleApi: systemModules.getPageRecycleList,
  rowSelection: {
    showCheckedAll: true
  }
})
const columns = reactive([
  {
    addDisplay: true,
    commonRules: {
      message: "请输入ID",
      required: false
    },
    dataIndex: "id",
    editDisplay: false,
    formType: "input",
    hide: false,
    search: true,
    sortable: {
      sortDirections: [
        "ascend",
        "descend"
      ],
      sorter: true
    },
    title: "ID"
  },
  {
    addDisplay: true,
    commonRules: {
      message: "请输入模块名称",
      required: true
    },
    dataIndex: "name",
    editDisplay: true,
    formType: "input",
    hide: false,
    search: true,
    sortable: {},
    title: "模块名称"
  },
  {
    addDisplay: true,
    commonRules: {
      message: "请输入模块标记",
      required: true
    },
    dataIndex: "label",
    editDisplay: true,
    formType: "input",
    hide: false,
    search: true,
    sortable: {},
    title: "模块标记"
  },
  {
    addDisplay: true,
    commonRules: {
      message: "请输入是否安装",
      required: true
    },
    dataIndex: "installed",
    dict: {
      name: "data_status",
      props: {
        label: "title",
        value: "key"
      },
      translation: true
    },
    editDisplay: true,
    formType: "radio",
    hide: false,
    search: true,
    sortable: {
      sortDirections: [
        "ascend",
        "descend"
      ],
      sorter: true
    },
    title: "是否安装"
  },
  {
    addDisplay: true,
    commonRules: {
      message: "请输入状态",
      required: true
    },
    dataIndex: "status",
    dict: {
      name: "data_status",
      props: {
        label: "title",
        value: "key"
      },
      translation: true
    },
    editDisplay: true,
    formType: "radio",
    hide: false,
    search: true,
    sortable: {
      sortDirections: [
        "ascend",
        "descend"
      ],
      sorter: true
    },
    title: "状态"
  },
  {
    addDisplay: false,
    commonRules: {
      message: "请输入创建时间",
      required: true
    },
    dataIndex: "created_at",
    editDisplay: false,
    formType: "time",
    hide: false,
    search: true,
    sortable: {},
    title: "创建时间"
  },
  {
    addDisplay: false,
    commonRules: {
      message: "请输入更新时间",
      required: true
    },
    dataIndex: "updated_at",
    editDisplay: false,
    formType: "input",
    hide: true,
    search: false,
    sortable: {},
    title: "更新时间"
  },
  {
    addDisplay: true,
    commonRules: {
      message: "请输入描述",
      required: true
    },
    dataIndex: "description",
    editDisplay: true,
    formType: "textarea",
    hide: false,
    search: false,
    sortable: {},
    title: "描述"
  },
  {
    addDisplay: false,
    commonRules: {
      message: "请输入创建者",
      required: true
    },
    dataIndex: "created_by",
    editDisplay: false,
    formType: "input",
    hide: true,
    search: false,
    sortable: {},
    title: "创建者"
  },
  {
    addDisplay: false,
    commonRules: {
      message: "请输入更新者",
      required: true
    },
    dataIndex: "updated_by",
    editDisplay: false,
    formType: "input",
    hide: true,
    search: false,
    sortable: {},
    title: "更新者"
  }
])
</script>
<script> export default {name: 'system:systemModules'} </script>
