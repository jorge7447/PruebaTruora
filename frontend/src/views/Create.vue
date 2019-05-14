<template>
    <div>
        <b-alert v-model="showError" variant="danger" dismissible>
            {{errorMessage}}
        </b-alert>
        <key-form :keyData="keyData" @processKey="addKey"></key-form>
    </div>
</template>

<script>
    import { mapState, mapActions } from 'vuex'
    import KeyForm from '@/components/KeyForm'
    export default {
        components: {
            KeyForm
        },
        data () {
            return {
                keyData: {
                    Name: ''
                },
                showError: false
            }
        },
        computed: {
            ...mapState('key', ['error', 'errorMessage'])
        },
        methods: {
            ...mapActions('key', ['create']),
            async addKey(keyData){
                this.showError = false
                await this.create(keyData)
                if(!this.error)
                    this.$router.push('/list')
                else{
                    this.showError = true
                }
            }
        }
    }
</script>
