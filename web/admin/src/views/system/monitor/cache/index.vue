<template>
  <a-layout-content>
    <div class="ma-content-block lg:flex p-4">
      <div class="flex justify-between w-full redis-info mt-3">
        <a-descriptions :column="2" size="large" bordered title="Redis信息" class="lg:w-10/12 w-full">
          <a-descriptions-item label="Redis版本">{{ server.version }}</a-descriptions-item>
          <a-descriptions-item label="客户端连接数">{{ server.clients }}</a-descriptions-item>
          <a-descriptions-item label="运行模式">{{ server.redis_mode }}</a-descriptions-item>
          <a-descriptions-item label="运行天数">{{ server.run_days }}</a-descriptions-item>
          <a-descriptions-item label="端口">{{ server.port }}</a-descriptions-item>
          <a-descriptions-item label="AOF状态">{{ server.aof_enabled }}</a-descriptions-item>
          <a-descriptions-item label="已过期key">{{ server.expired_keys }}</a-descriptions-item>
          <a-descriptions-item label="系统使用key">{{ server.sys_total_keys }}</a-descriptions-item>
        </a-descriptions>
        <div class="echarts hidden lg:block">
          <ma-chart :options="options" width="330px" height="330px" />
        </div>
      </div>
    </div>
    <div class="ma-content-block p-4 mt-3">
      <a-space>
        <a-link @click="clearAll()">清除所有缓存</a-link>
        <a-button
            v-if="selectedKeys.length > 0"
            @click="batchDelete"
            type="primary"
            status="danger"
            size="small"
        >
          批量删除 ({{ selectedKeys.length }})
        </a-button>
      </a-space>
    </div>
    <div class="ma-content-block p-4 mt-3">
      <div class="text-base">
        缓存数据管理
      </div>
      <div class="mt-5 lg:flex justify-between">
        <div class="lg:w-8/12 w-full">
          <a-input-search
              :model-value="searchKey"
              placeholder="输入关键词过滤缓存键"
              @input="handleSearch"
              class="mb-3"
          />
          <a-table
              :data="filteredData"
              :columns="columns"
              row-key="name"
              v-model:selected-keys="selectedKeys"
              :row-selection="{ type: 'checkbox', showCheckedAll: true }"
          >
            <template #operation="{ record }">
              <a-space>
                <a-link @click="viewKey(record.name)">查看</a-link>
                <a-popconfirm content="确实要删除该缓存吗?" position="bottom" @ok="del(record)">
                  <a-link>删除</a-link>
                </a-popconfirm>
              </a-space>
            </template>
          </a-table>
        </div>
        <a-textarea v-model="content" class="lg:w-4/12 w-full lg:ml-5 mt-3 lg:mt-0" readonly />
      </div>
    </div>
  </a-layout-content>
</template>

<script setup>
import { ref, reactive, computed,onUnmounted } from 'vue'
import monitor from '@/api/system/monitor'
import { Message, Modal } from '@arco-design/web-vue'
import { refreshTag } from '@/utils/common'

const searchKey = ref('')
const selectedKeys = ref([])
const options = ref({})
const server = ref({})
const data = ref([])
const content = ref('')

const columns = reactive([
  { title: '缓存键名', dataIndex: 'name' },
  { title: '操作', slotName: 'operation', width: 150, align: 'right' },
])

const filteredData = computed(() => {
  return data.value.filter(item =>
      item.name.toLowerCase().includes(searchKey.value.toLowerCase())
  )
})

// 新增防抖定时器
let timeoutId = null

// 修改后的handleSearch
const handleSearch = (value) => {
  clearTimeout(timeoutId)
  timeoutId = setTimeout(() => {
    searchKey.value = value
  }, 300)
}

// 在组件卸载时清除定时器
onUnmounted(() => {
  clearTimeout(timeoutId)
})


const viewKey = async (key) => {
  const response = await monitor.view({ key })
  content.value = response.data.content
}

const del = async (row) => {
  const response = await monitor.deleteKey({ key: row.name })
  if (response.success) {
    Message.success(response.message)
    await getCacheInfo()
    content.value = ''
  }
}

const clearAll = async () => {
  const response = await monitor.clear()
  if (response.success) {
    Message.success(response.message)
    await getCacheInfo()
    content.value = ''
  }
}

const batchDelete = () => {
    Modal.confirm({
      title: '确认删除',
      content: `确定要删除选中的 ${selectedKeys.value.length} 个缓存吗？`,
      onOk: () => {
        if (!confirm) return
        for (const key of selectedKeys.value) {
          monitor.deleteKey({ key })
        }
        Message.success('删除成功')
        selectedKeys.value = []
        getCacheInfo()
        content.value = ''
      }
    })
}

const getCacheInfo = async () => {
  const response = await monitor.getCacheInfo()
  server.value = response.data.server
  data.value = response.data.keys.map(item => ({ name: item }))

  options.value = {
    tooltip: {
      formatter: '{b} : {c} kb'
    },
    series: [
      {
        name: '内存占用情况',
        type: 'gauge',
        min: 0,
        max: 5 * 1024,
        progress: {
          show: true
        },
        detail: {
          valueAnimation: true,
          formatter: '{value}'
        },
        data: [
          {
            value: parseInt(response.data.server.use_memory),
            name: 'Redis占用内存'
          }
        ]
      }
    ]
  }
}

getCacheInfo()
</script>

<script>
export default { name: 'system:cache' }
</script>

<style scoped>
.redis-info {
  max-height: 260px;
  overflow: hidden;
}

.echarts {
  position: relative;
  top: -10px;
  right: -10px;
}

.arco-btn-danger {
  margin-left: 10px;
}
</style>