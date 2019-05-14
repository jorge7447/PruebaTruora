<template>

    <b-form @submit.prevent="$emit('processText', keyData.Text)">
        <b-form-group
            class="text-left font-weight-light"
            id="keyData"
            :label="labelTitle"
            label-for="keyData">

            <b-form-textarea
                autocomplete="off"
                id="keyData"
                v-model="keyData.Text"
                :state="!$v.keyData.Text.$invalid"
                placeholder="Introduce el texto"
                rows="3"
                @input="$v.keyData.$touch">
            </b-form-textarea>
            <b-form-invalid-feedback id="keyDataInfo" v-if="$v.keyData.$dirty">
               Este campo es requerido
            </b-form-invalid-feedback>
        </b-form-group>

        <b-button 
            type="submit"
            variant="outline-secondary"
            class="col-md-2" 
            :disabled="$v.keyData.$invalid">{{ detailSubmit }}</b-button>
    </b-form>
</template>

<script>
    import { validationMixin } from 'vuelidate'
    import { required, minLength } from 'vuelidate/lib/validators'
    export default {
        mixins: [validationMixin],
        props: {
            detailSubmit: {
                type: String,
                default: 'Encriptar'
            },
            labelTitle: {
                type: String,
                default: 'Texto a Encriptar'
            }
        },
        data () {
            return {
                keyData: {
                    Text: ''
                }
            }
        },
        validations: {
            keyData: {
                Text: {
                    required
                }
            }
        }
    }
</script>