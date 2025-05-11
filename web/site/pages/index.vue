<script setup>
import { Message } from '@arco-design/web-vue'
import { test } from '../api/test'

const isLoading = ref(false)
const helloWord = ref('')

async function handleExperienceClick() {
  try {
    isLoading.value = true
    const res = await test({})
    Message.success(`接口返回：${JSON.stringify(res)}`)
  }
  catch (error) {
    Message.error(`请求失败：${error.message}`)
  }
  finally {
    isLoading.value = false
  }
}

const { data: res } = await useAsyncData('test_key', () => test())
helloWord.value = res?.data

useSeoMeta({
  // 页面标题
  title: 'Devinggo - 基于 GoFrame v2 + Vue3 + Arco Design 开发的全栈前后端分离后台管理系统',
  // 页面描述
  description: 'Devinggo 是一个基于 GoFrame v2 + Vue3 + Arco Design 开发的全栈前后端分离的后台管理系统，提供高效、安全的开发框架。',
  // 页面关键词
  keywords: 'Devinggo, GoFrame, Vue3, Arco Design, 后台管理系统, 全栈开发, 前后端分离',
})
</script>

<template>
  <div class="min-h-screen dh-ct-bg">
    <!-- Hero Section -->
    <section class="bg-gradient-to-r from-blue-500 to-indigo-600 text-white py-20">
      <div class="container mx-auto text-center px-4">
        <h1 class="text-4xl md:text-6xl font-bold mb-4">
          欢迎来到 Devinggo
        </h1>
        <p class="text-lg md:text-xl mb-8">
          基于 GoFrame v2 + Vue3 + Arco Design 开发的全栈前后端分离后台管理系统
        </p>
        <NuxtLink to="https://devinggo-dev.devinghub.com/admin" class="bg-white text-indigo-600 font-semibold py-3 px-8 rounded-full shadow-lg hover:bg-gray-100 transition duration-300">
          开始探索
        </NuxtLink>
      </div>
    </section>

    <!-- Features Section -->
    <section class="py-16">
      <div class="container mx-auto px-4">
        <h2 class="text-3xl font-bold text-center mb-12 text-gray-800 dark:text-gray-200">
          核心特性
        </h2>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
          <!-- Feature 1 -->
          <div class="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-lg text-center">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 mx-auto mb-4 text-green-500" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
            </svg>
            <h3 class="text-xl font-semibold mb-2 text-gray-800 dark:text-gray-200">
              模块化设计
            </h3>
            <p class="text-gray-600 dark:text-gray-400">
              支持模块管理，代码生成，CRUD下载，提高开发效率。
            </p>
          </div>
          <!-- Feature 2 -->
          <div class="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-lg text-center">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 mx-auto mb-4 text-purple-500" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4" />
            </svg>
            <h3 class="text-xl font-semibold mb-2 text-gray-800 dark:text-gray-200">
              权限管理
            </h3>
            <p class="text-gray-600 dark:text-gray-400">
              完善的RBAC权限管理，支持多角色、多部门、数据权限控制。
            </p>
          </div>
          <!-- Feature 3 -->
          <div class="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-lg text-center">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 mx-auto mb-4 text-yellow-500" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
              <path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
            </svg>
            <h3 class="text-xl font-semibold mb-2 text-gray-800 dark:text-gray-200">
              多功能支持
            </h3>
            <p class="text-gray-600 dark:text-gray-400">
              支持多语言、多主题、队列、WebSocket、定时任务等丰富功能。
            </p>
          </div>
        </div>
      </div>
    </section>

    <!-- Call to Action Section -->
    <section class="bg-gray-100 dark:bg-gray-900 py-16">
      <div class="container mx-auto text-center px-4">
        <h2 class="text-3xl font-bold mb-4 text-gray-800 dark:text-gray-200">
          准备好提升您的开发效率了吗？
        </h2>
        <p class="text-lg text-gray-600 dark:text-gray-400 mb-8">
          {{ helloWord }}
        </p>
        <button
          class="bg-blue-500 hover:bg-blue-600 text-white font-semibold py-3 px-8 rounded-full shadow-lg transition duration-300"
          :disabled="isLoading"
          @click="handleExperienceClick"
        >
          {{ isLoading ? '请求中...' : '立即体验' }}
        </button>
      </div>
    </section>
  </div>
</template>

<style scoped>
/* 您可以在这里添加特定于此页面的样式 */
.dh-ct-bg {
  /* 保持或调整背景色 */
}
</style>
