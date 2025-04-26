import axios from "@/plugins/axios";

export async function xsdValidate(file: File) {
  const formData = new FormData();
  formData.append("file", file);
  const response = await axios.post("upload/xsd", formData, {
    headers: {
      "Content-Type": "multipart/form-data",
    },
  });
  return response;
}

export async function rngValidate(file: File) {
  const formData = new FormData();
  formData.append("file", file);
  const response = await axios.post("upload/rng", formData, {
    headers: {
      "Content-Type": "multipart/form-data",
    },
  });
  return response;
}