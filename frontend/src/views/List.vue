<template>
    <div >
        <b-row>
            <b-col md="12" class="my-1">
                <b-input-group>
                    <b-form-input v-model="filter" placeholder="Buscar"></b-form-input>
                    <b-input-group-append>
                        <b-button :disabled="!filter" @click="filter = ''">Clear</b-button>
                    </b-input-group-append>
                </b-input-group>
            </b-col>
        </b-row>
        <br>
        <div v-if="keys.length">
            <b-table :items="keys" :fields="fields" 
                :sort-by.sync="sortBy"
                :sort-desc.sync="sortDesc" striped>

            <template slot="Acciones" slot-scope="row">
                <b-button pill
                    variant="dark"
                    @click="goToActions(row.item)"
                >Acciones</b-button>
            </template>
            
            </b-table>
        </div>

        <b-alert show variant="info" v-else>No hay llaves</b-alert>
    </div>
</template>

<script>
    import { mapActions, mapState, mapMutations } from 'vuex'
    export default {
        computed: {
            ...mapState('key', ['keys'])
        },
        data() {
            return {
                sortBy: 'Name',
                sortDesc: false,
                filter: null,
                fields: [
                    { key: 'Name', sortable: true },
                    { key: 'CreatedAt', sortable: true },
                    { key: 'Acciones', sortable: false }
                ],
            }
        },
        methods: {
            ...mapActions('key', ['getKeys']),
            ...mapMutations('key', ['setKey']),
            goToActions(item){
                this.$router.push({name: 'detail', params: {id: item.ID}})
            }
        },
        mounted() {
           this.getKeys()
        },
        watch: {
            filter(search) {
                this.getKeys(search)
            }
        }
    }
</script>