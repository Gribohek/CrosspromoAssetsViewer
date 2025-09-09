<template>
  <v-card>
    <v-card-title class="d-flex align-center">
      <span class="text-h5">Crosspromo Assets</span>
      <v-spacer></v-spacer>
      <v-select
        v-model="selectedStoreCode"
        :items="storeCodes"
        label="Store Code"
        dense
        outlined
        class="mr-2"
        style="max-width: 200px"
        clearable
      ></v-select>
      <v-select
        v-model="selectedAppId"
        :items="appIds"
        label="App ID"
        dense
        outlined
        class="mr-2"
        style="max-width: 200px"
        clearable
      ></v-select>
      <v-btn color="primary" @click="fetchAssets" :loading="loading">
        Apply Filters
      </v-btn>
    </v-card-title>

    <v-card-text>
      <v-table density="comfortable">
        <thead>
          <tr>
            <th class="text-left">Localization</th>
            <th class="text-left">iOS</th>
            <th class="text-left">Android</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="assetGroup in assetGroups" :key="assetGroup.localization">
            <td>{{ assetGroup.localization }}</td>
            <td>
              <div class="d-flex flex-column">
                <AssetCell
                  :assets="assetGroup.platforms?.iOS?.banner"
                  type="Banner"
                  platform="iOS"
                />
                <AssetCell
                  :assets="assetGroup.platforms?.iOS?.interstitial"
                  type="Interstitial"
                  platform="iOS"
                />
                <AssetCell
                  :assets="assetGroup.platforms?.iOS?.rewarded"
                  type="Rewarded"
                  platform="iOS"
                />
              </div>
            </td>
            <td>
              <div class="d-flex flex-column">
                <AssetCell
                  :assets="assetGroup.platforms?.Android?.banner"
                  type="Banner"
                  platform="Android"
                />
                <AssetCell
                  :assets="assetGroup.platforms?.Android?.interstitial"
                  type="Interstitial"
                  platform="Android"
                />
                <AssetCell
                  :assets="assetGroup.platforms?.Android?.rewarded"
                  type="Rewarded"
                  platform="Android"
                />
              </div>
            </td>
          </tr>
        </tbody>
      </v-table>
    </v-card-text>

    <AssetModal
      v-model="modalVisible"
      :asset="selectedAsset"
      :platform="selectedPlatform"
      :type="selectedType"
    />
  </v-card>
</template>

<script>
import { api } from "../api";
import AssetCell from "./AssetCell.vue";
import AssetModal from "./AssetModal.vue";

export default {
  name: "AssetTable",
  components: {
    AssetCell,
    AssetModal,
  },
  data() {
    return {
      assetGroups: [],
      storeCodes: [],
      appIds: [],
      selectedStoreCode: null,
      selectedAppId: null,
      loading: false,
      modalVisible: false,
      selectedAsset: null,
      selectedPlatform: "",
      selectedType: "",
    };
  },
  async mounted() {
    await this.fetchStoreCodes();
    await this.fetchAppIds();
    await this.fetchAssets();
  },
  methods: {
    async fetchAssets() {
      this.loading = true;
      try {
        const filters = {
          storeCode: this.selectedStoreCode,
          appId: this.selectedAppId,
        };
        this.assetGroups = await api.getAssets(filters);
      } catch (error) {
        console.error("Error fetching assets:", error);
        this.assetGroups = [];
      } finally {
        this.loading = false;
      }
    },
    async fetchStoreCodes() {
      try {
        this.storeCodes = await api.getStoreCodes();
      } catch (error) {
        console.error("Error fetching store codes:", error);
        this.storeCodes = [];
      }
    },
    async fetchAppIds() {
      try {
        this.appIds = await api.getAppIds();
      } catch (error) {
        console.error("Error fetching app IDs:", error);
        this.appIds = [];
      }
    },
    openModal(asset, platform, type) {
      this.selectedAsset = asset;
      this.selectedPlatform = platform;
      this.selectedType = type;
      this.modalVisible = true;
    },
  },
};
</script>

<style scoped>
.v-table {
  background: transparent;
}
</style>
