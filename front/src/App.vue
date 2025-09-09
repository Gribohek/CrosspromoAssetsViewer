<template>
  <v-app>
    <v-main>
      <v-container>
        <AssetTable />
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import AssetTable from "./components/AssetTable.vue";

export default {
  name: "App",
  components: {
    AssetTable,
  },
};
</script>
