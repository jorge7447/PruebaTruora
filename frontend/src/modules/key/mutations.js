export function setKeys (state, keys) {
    state.keys = keys
}

export function setKey (state, selectedKey) {
    state.selectedKey = selectedKey
}

export function setError (state, payload) {
    state.error = payload.error
    if(payload.error){
        state.errorMessage = payload.message
    }else{
        state.errorMessage = ""
    }
}