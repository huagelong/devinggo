<template>
  <div class="ma-content-block lg:flex justify-between p-4">
    <!-- CRUD 组件 -->
    <ma-crud :options="crud" :columns="columns" ref="crudRef">
      <template #operationCell="{ record }">
        <a-popconfirm content="确实要将该用户强制退出吗?" position="bottom" @ok="kick(record)">
          <a-link v-auth="['system:onlineUser:kick']"><icon-import /> 强制退出</a-link>
        </a-popconfirm>
      </template>
    </ma-crud>
  </div>
</template>

<script setup>
  import { ref, reactive } from 'vue'
  import monitor from '@/api/system/monitor'
  import { Message } from '@arco-design/web-vue'

  const crudRef = ref()

  const kick = async (row) => {
    const response = await monitor.kickUser({ id: row.id,app_id:row.app_id })
    response.success && Message.success(response.message)
  }

  const crud = reactive({
    api: monitor.getOnlineUserPageList,
    showIndex: false,
    operationColumn: true,
    operationColumnWidth: 120,
    searchColNumber: 2,
    pageLayout: 'fixed',
  })

  const columns = reactive([
    { title: '用户账户', dataIndex: 'username', search: true, width: 180 },
    { title: '用户昵称', dataIndex: 'nickname', width: 180 },
    { title: 'App', dataIndex: 'app_id', width: 180,
      search: true,
      formType: 'select',
      dict: {
        remote: '/system/app/remote',
        translation:true,
        props: { label: 'app_name', value: 'app_id' },
        openPage: true,
        // 远程请求配置项
        remoteOption: {
          // 按用户名排序
          sort: { id: 'desc' }, // 如果不指定排序方式，默认为正序排序
          // 设置查询的字段
          select: [ 'app_id', 'app_name' ],
          // 设置数据过滤
          filter: {
          },
        }
      }
    },
    { title: '登录IP', dataIndex: 'login_ip', width: 180 },
    { title: '登录时间', dataIndex: 'login_time', width: 180 },
  ])
</script>

<script>
export default { name: 'system:onlineUser' }
</script>

<style scoped></style>