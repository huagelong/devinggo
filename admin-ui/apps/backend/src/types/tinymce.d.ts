declare module '@tinymce/tinymce-vue' {
  import type { DefineComponent } from 'vue';

  const Editor: DefineComponent<Record<string, unknown>, Record<string, unknown>, any>;
  export default Editor;
}

declare module 'tinymce' {
  const tinymce: any;
  export default tinymce;
}

declare module 'tinymce/*' {
  const mod: any;
  export default mod;
}
