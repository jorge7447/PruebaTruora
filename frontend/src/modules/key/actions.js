import Vue from 'vue'

export async function create({commit}, keyData) {
    try {
        commit('setError', {error: false})
        commit('setLoading', true, {root: true})
        const {data} = await Vue.axios({
            method: 'POST',
            url: '/v1/api/key',
            data: keyData
        })
        return data
    } catch (error) {
        commit('setError', {error: true, message: error.message+', verifique que el nombe de la llave no exista'})
    }finally{
        commit('setLoading',false, { root: true })
    }
}

export async function getKeys({commit}, search = null) {
    try {
        commit('setError', {error: false})
        commit('setLoading', true, {root: true})
        let searchUrl = ""
        if(search != null){
            searchUrl = '?search='+search
        }

        const {data} = await Vue.axios({
            url: '/v1/api/key'+searchUrl
        })
        commit('setKeys', data)
    } catch (error) {
        commit('setError', {error: true, message: error.message})
    }finally{
        commit('setLoading',false, { root: true })
    }
}

export async function getKey({ commit }, id) {
    try {
        commit('setError', {error: false})
        commit('setLoading', true, {root: true})
        const {data} = await Vue.axios({
            method: 'GET',
            url: `/v1/api/key/${id}`,
        })
        commit('setKey', data)   
    } catch (error) {
        commit('setError', {error: true, message: error.message})
    } finally {
        commit('setLoading',false, { root: true })
    }
}

export async function encrypt({commit}, keyData) {
    try {
        commit('setError', {error: false})
        commit('setLoading', true, {root: true})
        const {data} = await Vue.axios({
            method: 'POST',
            url: '/v1/api/key/encrypt',
            data: keyData
        })
        return data
    } catch (error) {
        commit('setError', {error: true, message: error.message})
    }finally{
        commit('setLoading',false, { root: true })
    }
}

export async function decrypt({commit}, keyData) {
    try {
        commit('setError', {error: false})
        commit('setLoading', true, {root: true})
        const {data} = await Vue.axios({
            method: 'POST',
            url: '/v1/api/key/decrypt',
            data: keyData
        })
        return data
    } catch (error) {
        commit('setError', {error: true, message: error.message})
    }finally{
        commit('setLoading',false, { root: true })
    }
}