<template>
  <div>
    <editor :key="editorKey" v-model="content" :init="initConfig" :id="props.id"></editor>

    <a-modal v-model:visible="resourceVisible" :width="1080" :footer="false" draggable>
      <template #title>资源选择器</template>
      <ma-resource v-model="list" multiple ref="resource" returnType="url" />
    </a-modal>
  </div>
</template>

<script setup>
import {reactive, ref, watch, computed} from 'vue'
import MaResource from '@cps/ma-resource/index.vue'
import {useAppStore} from '@/store'
import {uploadRequest} from "@cps/ma-upload/js/utils"
import uploadConfig from '@/config/upload'
import tool from '@/utils/tool'

import Editor from '@tinymce/tinymce-vue'
import tinymce from 'tinymce/tinymce'

import 'tinymce/icons/default'
import 'tinymce/models/dom'
import 'tinymce/themes/silver'

import 'tinymce/plugins/advlist'    //高级列表
import 'tinymce/plugins/anchor'   //锚点
import 'tinymce/plugins/autolink'   //自动链接
import 'tinymce/plugins/autosave'   //自动存稿
import 'tinymce/plugins/charmap'    //特殊字符
import 'tinymce/plugins/code'   //编辑源码
import 'tinymce/plugins/codesample'   //代码示例
import 'tinymce/plugins/directionality'   //文字方向
import 'tinymce/plugins/emoticons'    //表情
import 'tinymce/plugins/fullscreen'   //全屏
import 'tinymce/plugins/help'   //帮助
import 'tinymce/plugins/image'    //插入编辑图片
import 'tinymce/plugins/importcss'    //引入css
import 'tinymce/plugins/insertdatetime'   //插入日期时间
import 'tinymce/plugins/link'   //超链接
import 'tinymce/plugins/lists'    //列表插件
import 'tinymce/plugins/media'    //插入编辑媒体
import 'tinymce/plugins/nonbreaking'    //插入不间断空格
import 'tinymce/plugins/pagebreak'    //插入分页符
import 'tinymce/plugins/preview'    //预览
import 'tinymce/plugins/quickbars'    //快速工具栏
import 'tinymce/plugins/save'   //保存
import 'tinymce/plugins/searchreplace'    //查找替换
import 'tinymce/plugins/table'    //表格
import 'tinymce/plugins/template'   //内容模板
import 'tinymce/plugins/visualblocks'   //显示元素范围
import 'tinymce/plugins/visualchars'    //显示不可见字符
import 'tinymce/plugins/wordcount'    //字数统计


const appStore = useAppStore()

const props = defineProps({
  modelValue: {type: String},
  height: {type: Number, default: 400},
  id: {type: String, default: () => 'tinymce' + new Date().getTime().toString()},
  plugins: {
    type: [String, Array],
    default:
        "preview searchreplace autolink directionality visualblocks visualchars fullscreen link media code codesample table charmap nonbreaking insertdatetime advlist lists wordcount autosave"
  },
  toolbar: {
    type: [String, Array],
    default:
        "code undo redo restoredraft | paste pastetext | forecolor backcolor bold italic underline strikethrough link codesample | fullscreen preview | alignleft aligncenter alignright alignjustify outdent indent formatpainter | \
      styles codeformat fontfamily fontsize | bullist numlist | blockquote subscript superscript removeformat | table tablecellborderstyle media \
      charmap pagebreak insertdatetime | resource"
  }
})

const emit = defineEmits(['update:modelValue', 'change'])

let content = computed({
  get() {
    return props.modelValue
  },
  set(value) {
    emit('update:modelValue', value)
  },
});

const list = ref([])
const resource = ref()
const resourceVisible = ref(false)
const appBaseUrl = import.meta.env.VITE_APP_BASE

const initConfig = reactive({
  menubar: false, // 菜单栏显隐
  language_url: `${appBaseUrl}tinymce/i18n/zh_CN.js`,
  language: 'zh_CN',
  skin_url: appBaseUrl + (appStore.mode === 'light' ? 'tinymce/skins/ui/tinymce-5' : 'tinymce/skins/ui/tinymce-5-dark'),
  height: props.height,
  toolbar_mode: 'wrap',
  plugins: props.plugins,
  toolbar: props.toolbar,
  branding: false,
  content_css: `${appBaseUrl}tinymce/skins/content/default/content.css`,
  font_size_formats: '12px 14px 16px 18px 20px 22px 24px 26px 28px 30px 32px 34px 36px 38px 40px',
  font_size_input_default_unit: 'px',
  images_upload_handler: async (blobInfo, success) => {
    let res = await uploadRequest(blobInfo.blob(), 'image', 'uploadImage')
    if (!res) throw ('上传错误');
    return /image/g.test(res.mime_type) ? tool.attachUrl(res.url, getStoreMode(res.storage_mode)) : $url + 'not-image.png'
  },
  setup: (editor) => {
    editor.on('init', () => {
      editor.getBody().style.fontSize = '14px';
    })
    editor.ui.registry.addButton('resource', {
      text: '资源选择器',
      onAction: () => resourceVisible.value = true
    })
  }
})

const getStoreMode = (mode) => {
  return uploadConfig.storageMode[mode.toString()]
}

const editorKey = ref(new Date().getTime())

watch(
    () => list.value,
    imgs => {
      let tmp = ''
      imgs.map(img => {
        if (img.indexOf('.jpg') > -1 || img.indexOf('.png') > -1 || img.indexOf('.bmp') > -1 || img.indexOf('.jpeg') > -1 || img.indexOf('.svg') > -1 || img.indexOf('.gif') > -1) {
          tmp += `<img src=${img} width="100%" />`
        }
      })
      content.value = content.value ? content.value + tmp : tmp
      resource.value.clearSelecteds()
      resourceVisible.value = false
    }
)
watch(
    () => content.value,
    vl => emit('change', vl)
)
</script>
