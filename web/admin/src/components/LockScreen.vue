<template>
  <div class="lock-screen" v-if="visible">
    <div class="lock-screen-content">
      <a-row :gutter="24">
        <a-col :span="18">
          <a-input v-model="password" type="password" placeholder="请输入锁屏密码" />
        </a-col>
        <a-col :span="6">
          <a-button type="primary" @click="unlock">解锁</a-button>
        </a-col>
      </a-row>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue'
import tool from '@/utils/tool';
import { Message } from '@arco-design/web-vue';
const emit = defineEmits(['update:visible']);
import { useAppStore } from '@/store'
const appStore = useAppStore()
const lockScreenPwd = appStore.getLockScreenPwd();

// 修改 props 的定义方式，将其赋值给一个变量
const props = defineProps({
  visible: {
    type: Boolean,
    required: true,
    default: false,
  }
});

const password = ref('');

const unlock = () => {
  if (tool.md5(password.value) === lockScreenPwd) {
    password.value = '';
    emit('update:visible', false);
    appStore.setIsLocked(false);
  } else {
    Message.error('锁屏密码错误');
  }
};

const lockScreen = () => {
  appStore.setIsLocked(true);
  emit('update:visible', true); 
};

const handleKeydown = (event) => {
  // 使用 props.visible 来访问 visible prop
  if (!props.visible) return;
  
  if (event.altKey && event.key === 'n') {
    lockScreen();
  }
  if (event.key === 'Enter') {
    event.preventDefault();
    unlock();
  }
};

onMounted(() => {
  window.addEventListener('keydown', handleKeydown);
});

onBeforeUnmount(() => {
  window.removeEventListener('keydown', handleKeydown);
});

</script>

<style scoped>
.lock-screen {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: #999;
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.lock-screen-content {
  background: white;
  padding: 20px;
  border-radius: 5px;
}
</style>
