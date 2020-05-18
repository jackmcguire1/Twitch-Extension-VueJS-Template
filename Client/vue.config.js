module.exports = {
  configureWebpack: {
    devtool: 'source-map',
  },

  chainWebpack: config => {
    config.externals({
      ...config.get('externals'),
      'ga': 'ga',
      'twitch': 'Twitch'
    })
  },

  pages: {
    panel: {
      entry: './src/main.js',
      template: './twitch-template.html',
      filename: 'panel.html',
      title: 'twitch-ext-vuejs-tmpl'
    },
    config: {
      entry: './src/main.js',
      template: './twitch-template.html',
      filename: 'config.html',
      title: 'twitch-ext-vuejs-tmpl',
    },
    liveconfig: {
      entry: './src/main.js',
      template: './twitch-template.html',
      filename: 'liveconfig.html',
      title: 'twitch-ext-vuejs-tmpl'
    }
  },

  runtimeCompiler: true,

  devServer: {
    clientLogLevel: 'info',
    proxy: {
      '/api': {
        target: 'https://localhost:8008',
        changeOrigin: true,
        pathRewrite: {
          '^/api': ''
        }
      }
    },
    watchOptions: {
      poll: true
    },
    disableHostCheck: true
  },

  publicPath: '',
  outputDir: 'dist',
  assetsDir: 'static'
}