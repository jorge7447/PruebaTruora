import Vue from 'vue'
import Vuex from 'vuex'

import key from '@/modules/key'
Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    loading: false
  },
  mutations: {
    setLoading(state, loading) {
      state.loading = loading
    }
  },
  modules: {
    key
  }
})
