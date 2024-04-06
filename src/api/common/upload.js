import config from '@/config';
import http from '@/utils/request';

export default {
  upload: {
    url: `${config.API_URL}/admin/common/upload`,
    name: '文件上传',
    post: async function (params) {
      return await http.post(this.url, params);
    }
  }
};
