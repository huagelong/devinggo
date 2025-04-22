import Mock from 'mockjs'

export default defineEventHandler(() => {
  const mockData = {
    'list|18': [
      {
        id: '@id(6)',
        name: '@name',
        title: '@ctitle',
        createTime: '@date',
        updateTime: '@date',
        state: '@boolean',
        remark: '@cparagraph(1)',
      },
    ],
    'total': 18,
  }
  const rs = formatMockResult(200, mockData, '')
  return Mock.mock(rs)
})
