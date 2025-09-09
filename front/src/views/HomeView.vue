<template>
  <div class="home">
    <AssetTable />
  </div>
</template>

<script>
import AssetTable from "../components/AssetTable.vue";

export default {
  name: "HomeView",
  components: {
    AssetTable,
  },
};
</script>

<style>
.home {
  padding: 20px;
}
</style>
