<script lang="ts" setup>
import { logger } from '#/utils/logger';
import type { EchartsUIType } from '@vben/plugins/echarts';

import { onMounted, ref } from 'vue';

import { EchartsUI, useEcharts } from '@vben/plugins/echarts';

import { getDashboardLoginChartApi } from '#/api/core/dashboard';
import { $t } from '#/locales';

const chartRef = ref<EchartsUIType>();
const { renderEcharts } = useEcharts(chartRef);

onMounted(async () => {
  try {
    // 默认获取过去10天的登录数据
    const data = await getDashboardLoginChartApi(10);

    renderEcharts({
      tooltip: {
        trigger: 'axis',
        axisPointer: {
          lineStyle: {
            color: '#5ab1ef',
            width: 1,
          },
        },
      },
      grid: {
        bottom: 0,
        containLabel: true,
        left: '1%',
        right: '1%',
        top: '10%',
      },
      xAxis: {
        type: 'category',
        boundaryGap: false,
        data: data.xAxis || [],
        axisTick: {
          show: false,
        },
        splitLine: {
          show: true,
          lineStyle: {
            type: 'solid',
            width: 1,
            color: 'rgba(0,0,0,0.05)',
          },
        },
      },
      yAxis: [
        {
          type: 'value',
          axisTick: {
            show: false,
          },
          splitArea: {
            show: true,
          },
          splitNumber: 4,
          splitLine: {
            lineStyle: {
              color: 'rgba(0,0,0,0.05)',
            },
          },
        },
      ],
      series: [
        {
          name: $t('dashboard.analytics.loginTimes'),
          type: 'line',
          smooth: true,
          data: data.chartsData || [],
          areaStyle: {
            opacity: 0.1,
          },
          itemStyle: {
            color: '#5ab1ef',
          },
        },
      ],
    });
  } catch (error) {
    logger.error('Failed to load dashboard login chart', error);
  }
});
</script>

<template>
  <EchartsUI ref="chartRef" />
</template>
