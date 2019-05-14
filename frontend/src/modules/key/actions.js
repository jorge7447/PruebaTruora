import Vue from 'vue'

export async function create({commit}, keyData) {
    try {
        commit('setLoading', true, {root: true})
        const {data} = await Vue.axios({
            method: 'POST',
            url: '/v1/api/key',
            data: keyData
        })
        return data
    } catch (error) {
        commit('key/keyError', error.message, { root: true })
    }finally{
        commit('setLoading',false, { root: true })
    }
}

export async function getKeys({commit}, search = null) {
    try {
        
        let searchUrl = ""
        if(search != null){
            commit('setLoading', true, {root: true})
            searchUrl = '?search='+search
        }

        const {data} = await Vue.axios({
            url: '/v1/api/key'+searchUrl
        })
        //alert(data[0].ID)
        commit('setKeys', data)
    } catch (error) {
        commit('key/keyError', error.message, { root: true })
    }finally{
        commit('setLoading',false, { root: true })
    }
}

export async function getKey({ commit }, id) {
    try {
        commit('setLoading', true, {root: true})
        const {data} = await Vue.axios({
            method: 'GET',
            url: `/v1/api/key/${id}`,
        })
        commit('setKey', data)   
    } catch (error) {
        commit('key/keyError', error.message, { root: true })
    } finally {
        commit('setLoading',false, { root: true })
    }
}

export async function encrypt({commit}, keyData) {
    try {
        commit('setLoading', true, {root: true})
        const {data} = await Vue.axios({
            method: 'POST',
            url: '/v1/api/key/encrypt',
            data: keyData
        })
        return data
    } catch (error) {
        commit('key/keyError', error.message, { root: true })
    }finally{
        commit('setLoading',false, { root: true })
    }
}

export async function decrypt({commit}, keyData) {
    try {
        commit('setLoading', true, {root: true})
        const {data} = await Vue.axios({
            method: 'POST',
            url: '/v1/api/key/decrypt',
            data: keyData
        })
        return data
    } catch (error) {
        commit('key/keyError', error.message, { root: true })
    }finally{
        commit('setLoading',false, { root: true })
    }
}