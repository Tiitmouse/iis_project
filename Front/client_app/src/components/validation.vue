
<template>
  <div>
    <H3>Validate with:</H3>
    <v-radio-group v-model="validationType" inline>
      <v-radio label="XSD" value="xsd"></v-radio>
      <v-radio label="RelaxNG" value="rng"></v-radio>
    </v-radio-group>
  </div>
  <div>
    <v-file-upload
      clearable
      density="compact"
      title="Drag and drop .xml file"
      variant="compact"
      accept=".xml"
      @change="handleFileUpload"
    ></v-file-upload>
  </div>
  <div>
    <v-btn class="btnValidate" rounded="xs" block @click="validateFile">
      Validate
    </v-btn>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { VFileUpload } from "vuetify/labs/VFileUpload";

const validationType = ref("xsd");
const uploadedFilePath = ref<string | undefined>();
    
function handleFileUpload(event: Event) {
  const input = event.target as HTMLInputElement;
  if (input.files && input.files[0]) {
    uploadedFilePath.value = input.files[0].name;
  }
  console.log(uploadedFilePath);
}

async function validateFile() {
  if (!uploadedFilePath.value) {
    alert("Please upload a file.");
    return;
  }

  if (validationType.value === "xsd") {
    // TODO: xsd
    alert("XSD validation");
  } else if (validationType.value === "rng") {
    // TODO: rng
    alert("RelaxNG validation");
  }
}

</script>

<style>
.v-sheet {
  background: #5373b3;
}
.btnValidate{
  background-color: white;
  color: #5373b3;
  font-weight: bold;
  margin-top: 20px;
  box-shadow:  0 0 15px 2px #5373b3;
}
</style>
