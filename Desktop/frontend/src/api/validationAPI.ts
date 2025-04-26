import { Validate } from "../../wailsjs/go/main/App"

export async function xsdValidate(file: File) {
  //@ts-ignore
  const dataBytes = await file.bytes();
  const decoder = new TextDecoder('utf-8')
  const data = decoder.decode(dataBytes);
  const rez = Validate({ Name: file.name, Size: file.size, Type: file.type }, data, "xsd")
  return rez
}

export async function rngValidate(file: File) {
  //@ts-ignore
  const dataBytes = await file.bytes();
  const decoder = new TextDecoder('utf-8')
  const data = decoder.decode(dataBytes);
  const rez = Validate({ Name: file.name, Size: file.size, Type: file.type }, data, "rng")
  return rez
}