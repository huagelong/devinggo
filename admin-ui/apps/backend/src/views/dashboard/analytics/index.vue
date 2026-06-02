<script lang="ts" setup>
import { logger } from '#/utils/logger';
import type { AnalysisOverviewItem } from '@vben/common-ui';
import type { TabOption } from '@vben/types';

import { markRaw, onMounted, ref } from 'vue';

import { AnalysisChartsTabs, AnalysisOverview } from '@vben/common-ui';
import {
  SvgBellIcon,
  SvgCakeIcon,
  SvgCardIcon,
  SvgDownloadIcon,
} from '@vben/icons';

import { getDashboardStatisticsApi } from '#/api/core/dashboard';
import { $t } from '#/locales';

import AnalyticsTrends from './analytics-trends.vue';

const overviewItems = ref<AnalysisOverviewItem[]>([
  {
    icon: markRaw(SvgCardIcon),
    title: $t('dashboard.analytics.totalUsers'),
    totalTitle: $t('dashboard.analytics.newUsers'),
    totalValue: 0,
    value: 0,
  },
  {
    icon: markRaw(SvgDownloadIcon),
    title: $t('dashboard.analytics.totalAttachments'),
    totalTitle: $t('dashboard.analytics.newAttachments'),
    totalValue: 0,
    value: 0,
  },
  {
    icon: markRaw(SvgCakeIcon),
    title: $t('dashboard.analytics.totalLogins'),
    totalTitle: $t('dashboard.analytics.newLogins'),
    totalValue: 0,
    value: 0,
  },
  {
    icon: markRaw(SvgBellIcon),
    title: $t('dashboard.analytics.totalOperations'),
    totalTitle: $t('dashboard.analytics.newOperations'),
    totalValue: 0,
    value: 0,
  },
]);

const chartTabs: TabOption[] = [
  {
    label: $t('dashboard.analytics.loginChart'),
    value: 'trends',
  },
];

async function initData() {
  try {
    const data = await getDashboardStatisticsApi();
    overviewItems.value = [
      {
        icon: markRaw(SvgCardIcon),
        title: $t('dashboard.analytics.userCount'),
        totalTitle: $t('dashboard.analytics.totalUserCount'),
        totalValue: data.userStats?.total || 0,
        value: data.userStats?.new || 0,
      },
      {
        icon: markRaw(SvgDownloadIcon),
        title: $t('dashboard.analytics.attachmentCount'),
        totalTitle: $t('dashboard.analytics.totalAttachmentCount'),
        totalValue: data.attachmentStats?.total || 0,
        value: data.attachmentStats?.new || 0,
      },
      {
        icon: markRaw(SvgCakeIcon),
        title: $t('dashboard.analytics.loginCount'),
        totalTitle: $t('dashboard.analytics.totalLoginCount'),
        totalValue: data.loginStats?.total || 0,
        value: data.loginStats?.new || 0,
      },
      {
        icon: markRaw(SvgBellIcon),
        title: $t('dashboard.analytics.operationCount'),
        totalTitle: $t('dashboard.analytics.totalOperationCount'),
        totalValue: data.operationStats?.total || 0,
        value: data.operationStats?.new || 0,
      },
    ];
  } catch (error) {
    logger.error('Failed to load dashboard statistics', error);
  }
}

onMounted(() => {
  initData();
});
</script>

<template>
  <div class="p-5">
    <AnalysisOverview :items="overviewItems" />
    <AnalysisChartsTabs :tabs="chartTabs" class="mt-5">
      <template #trends>
        <AnalyticsTrends />
      </template>
    </AnalysisChartsTabs>
  </div>
</template>
