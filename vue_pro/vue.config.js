module.exports = {
  transpileDependencies: [
    'vuetify'
  ],
  devServer: {
    hot: true,
    proxy: {
      "/api": {
        // target: "https://lives-api.chungoulife.com/Api",
        // target: 'https://gray-lives-api.chungoulife.com/Api',
        // target: 'https://ready-lives-api.chungoulife.com/Api',
        // target: 'http://testnewapi.chungoulife.com/Api',
        target: 'http://localhost:9000/api',
        changeOrigin: true,
        pathRewrite: {
          "^/api": ""
        }
      }
    },
  }
}
