module.exports = {
  entry: "./front/index.js",
  output: {
    path: __dirname,
    filename: "./public/js/index.bundle.js"
  },
  resolve: {
    alias: {
      vue: "vue/dist/vue.esm.js",
      vuex: "vuex/dist/vuex.js"
    }
  },
  module: {
    rules: [
      {
        test: /\.vue$/,
        loader: "vue-loader"
      }
    ]
  }
};
