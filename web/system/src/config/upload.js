const env = import.meta.env;
export default {
  storage: {
    LOCAL: env.VITE_UPLOAD_LOCAL_URL,
    OSS: env.VITE_UPLOAD_OSS_URL,
    QINIU: env.VITE_UPLOAD_QINIU_URL,
    COS: env.VITE_UPLOAD_COS_URL,
    HUAWEICLOUD: env.VITE_UPLOAD_HUAWEICLOUD_URL,
    AWSS3: env.VITE_UPLOAD_AWSS3_URL,
    MINIO: env.VITE_UPLOAD_MINIO_URL,
  },

  storageMode: {
    1: 'LOCAL',
    2: 'OSS',
    3: 'QINIU',
    4: 'COS',
    5: 'HUAWEICLOUD',
    6: 'AWSS3',
    7: 'MINIO',
  },
};
