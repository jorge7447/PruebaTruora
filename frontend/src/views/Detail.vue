<template>
    <div v-if="selectedKey">
        <p class="h4 p-1">{{selectedKey.Name}}, {{selectedKey.CreatedAt.slice(0,10)}}</p>

        <b-card border-variant="light" header="Encriptar" class="text-center">
            <detail-form @processText="encryptText"></detail-form>
            <br>
            <b-form-textarea v-if="cipher.show"
                v-model="cipher.message"
                rows="4"
                readonly>
            </b-form-textarea>
        </b-card>
        <br><br>
        <b-card border-variant="light" header="Desencriptar" class="text-center">
            <detail-form detailSubmit="Desencriptar" labelTitle="Texto a Desencriptar" @processText="decryptText"></detail-form>
            <br>            
            <b-form-textarea v-if="plain.show"
                v-model="plain.message"
                rows="2"
                readonly>
            </b-form-textarea>
        </b-card>
    </div>
</template>

<script>
    import { mapState, mapActions } from 'vuex'
    import DetailForm from '@/components/DetailForm'
    export default {
        components: {
            DetailForm
        },
        data () {
            return {
                cipher: {
                    show: false,
                    message: ""
                },
                plain: {
                    show: false,
                    message: ""
                },
            }
        },
        computed: {
            ...mapState('key', ['error', 'errorMessage', 'selectedKey'])
        },
        methods: {
            ...mapActions('key', ['getKey', 'encrypt', 'decrypt']),
            async encryptText(plainText){
               let result = await this.encrypt({"Id": this.selectedKey.ID, "Message": plainText})
               this.cipher.message = result.message
               this.cipher.show = true
            },
            async decryptText(cipherText){

               let result = await this.decrypt({"Id": this.selectedKey.ID, "Message": cipherText})
               this.plain.message = result.message
               this.plain.show = true
            }
        },
        mounted() {
            this.getKey(this.$route.params.id)
        }
    }
</script>