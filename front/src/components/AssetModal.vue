<template>
  <div v-if="modelValue" class="modal-overlay" @click.self="closeModal">
    <div class="modal">
      <div class="modal-header">
        <h3>{{ modalTitle }}</h3>
        <button class="modal-close" @click="closeModal">×</button>
      </div>

      <div class="modal-content">
        <!-- Превью изображения -->
        <div v-if="isImage" class="preview-container">
          <img :src="asset.url" :alt="asset.fileName" class="preview-image" />
        </div>

        <!-- Превью видео -->
        <div v-else-if="isVideo" class="preview-container">
          <video :src="asset.url" controls class="preview-video">
            Ваш браузер не поддерживает видео
          </video>
        </div>

        <!-- Информация об файле -->
        <div class="asset-info">
          <p><strong>Имя файла:</strong> {{ asset.fileName }}</p>
          <p><strong>Размер:</strong> {{ formatFileSize(asset.fileSize) }}</p>
          <p><strong>Тип:</strong> {{ asset.mimeType }}</p>
        </div>
      </div>

      <div class="modal-footer">
        <button @click="closeModal" class="close-button">Закрыть</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "AssetModal",
  props: {
    // Переименовали modelValue в isVisible для понятности
    isVisible: {
      type: Boolean,
      required: true,
    },
    asset: {
      type: Object,
      default: () => ({}),
    },
    platform: {
      type: String,
      default: "",
    },
    type: {
      type: String,
      default: "",
    },
  },
  emits: ["update:visible"], // Правильное объявление события
  computed: {
    modalTitle() {
      return `${this.platform} - ${this.type}`;
    },
    isImage() {
      return this.asset?.mimeType?.includes("image");
    },
    isVideo() {
      return this.asset?.mimeType?.includes("video");
    },
  },
  methods: {
    formatFileSize(size) {
      if (size < 1) {
        return (size * 1024).toFixed(1) + " KB";
      }
      return size.toFixed(1) + " MB";
    },
    closeModal() {
      // Вместо изменения props, вызываем событие
      this.$emit("update:visible", false);
    },
  },
};
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal {
  background: white;
  border-radius: 8px;
  padding: 20px;
  max-width: 500px;
  width: 90%;
  max-height: 80vh;
  overflow-y: auto;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 15px;
  border-bottom: 1px solid #eee;
  padding-bottom: 10px;
}

.modal-close {
  background: none;
  border: none;
  font-size: 24px;
  cursor: pointer;
  color: #666;
}

.modal-content {
  text-align: center;
}

.preview-image {
  max-width: 100%;
  max-height: 300px;
  border-radius: 4px;
}

.preview-video {
  max-width: 100%;
  max-height: 300px;
  border-radius: 4px;
}

.preview-placeholder {
  padding: 40px;
  text-align: center;
  color: #666;
  background: #f9f9f9;
  border-radius: 4px;
  margin-bottom: 15px;
}

.asset-info {
  background: #f9f9f9;
  padding: 15px;
  border-radius: 4px;
  margin-top: 15px;
  text-align: left;
}

.asset-info p {
  margin: 5px 0;
}

.modal-footer {
  margin-top: 20px;
  text-align: center;
}

.close-button {
  padding: 10px 20px;
  background: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.close-button:hover {
  background: #0056b3;
}
</style>
