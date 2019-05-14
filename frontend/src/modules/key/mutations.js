export function setKeys (state, keys) {
    state.keys = keys
    console.log("setKey", state.keys)
}

export function setKey (state, selectedKey) {
    state.selectedKey = selectedKey
}

export function keyError (state, payload) {
    state.error = true
    state.errorMessage = payload
    //state.keys = []
}