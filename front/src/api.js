import axios from "axios";

const API_BASE_URL = "http://localhost:3000/api";

export const api = {
  // Get assets with optional filters
  async getAssets(filters = {}) {
    const params = new URLSearchParams();
    Object.keys(filters).forEach((key) => {
      if (filters[key]) {
        params.append(key, filters[key]);
      }
    });

    const response = await axios.get(`${API_BASE_URL}/assets?${params}`);
    return response.data;
  },

  // Get available store codes
  async getStoreCodes() {
    const response = await axios.get(`${API_BASE_URL}/store-codes`);
    return response.data;
  },

  // Get available app IDs
  async getAppIds() {
    const response = await axios.get(`${API_BASE_URL}/app-ids`);
    return response.data;
  },
};
