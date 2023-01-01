<template>
  <div>{{title}}</div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  name: 'Tutorial',
  data: () => {
    return {
      title : 'titleee'
    }
  },
  mounted(){
    console.log('mounted', this.$store.getters.getCounter)
    console.log('mounted', this.$store)
    this.$store.dispatch('login', 'hai')
  }
})
</script>
