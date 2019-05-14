<template>
    <b-form @submit.prevent="$emit('processKey', keyData)">
        <b-card border-variant="light" header="Crear Llave RSA" class="text-center">
            <b-form-group
                class="text-left font-weight-light"
                id="keyData"
                label="Nombre"
                label-for="keyData">

                <b-form-input 
                    autocomplete="off"
                    id="keyData"
                    v-model="keyData.Name"
                    :state="!$v.keyData.Name.$invalid"
                    placeholder="Introduce el nombre"
                    @input="$v.keyData.$touch">
                </b-form-input>
                <b-form-invalid-feedback id="keyDataInfo" v-if="$v.keyData.$dirty">
                Este campo es requerido y debe tener una longitud minima de 4 
                </b-form-invalid-feedback>
            </b-form-group>

            <b-button 
                type="submit"
                variant="outline-secondary"
                class="col-md-2" 
                :disabled="$v.keyData.$invalid">Crear</b-button>
        </b-card>

    </b-form>
</template>

<script>
    import { validationMixin } from 'vuelidate'
    import { required, minLength } from 'vuelidate/lib/validators'
    export default {
        mixins: [validationMixin],
        props: {
            keyData: {
                type: Object,
                required: true
            }
        },
        validations: {
            keyData: {
                Name: {
                    required, minLength: minLength(4)
                }
            }
        }
    }
</script>