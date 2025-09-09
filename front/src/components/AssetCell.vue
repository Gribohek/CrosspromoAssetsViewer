<template>
  <div class="asset-cell">
    <template v-if="assets && assets.length > 0">
      <template v-if="assets.length === 1">
        <v-btn
          variant="text"
          color="primary"
          @click="$emit('open-modal', assets[0], platform, type)"
          class="text-none"
        >
          {{ formatFileSize(assets[0].fileSize) }}
        </v-btn>
      </template>
      <template v-else>
        <v-menu>
          <template v-slot:activator="{ props }">
            <v-btn
              variant="text"
              color="primary"
              v-bind="props"
              class="text-none"
            >
              {{ assets.length }} files
            </v-btn>
          </template>
          <v-list>
            <v-list-item
              v-for="(asset, index) in assets"
              :key="index"
              @click="$emit('open-modal', asset, platform, type)"
            >
              <v-list-item-title>
                {{ formatFileSize(asset.fileSize) }}
              </v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
      </template>
    </template>
    <template v-else>
      <span class="text-grey">-</span>
    </template>
  </div>
</template>

<script>
export default {
  name: "AssetCell",
  props: {
    assets: {
      type: Array,
      default: () => [],
    },
    type: {
      type: String,
      required: true,
    },
    platform: {
      type: String,
      required: true,
    },
  },
  emits: ["open-modal"],
  methods: {
    formatFileSize(size) {
      if (size < 1) {
        return `${(size * 1024).toFixed(1)} KB`;
      }
      return `${size.toFixed(1)} MB`;
    },
  },
};
</script>

<style scoped>
.asset-cell {
  min-height: 40px;
  display: flex;
  align-items: center;
}
</style>
