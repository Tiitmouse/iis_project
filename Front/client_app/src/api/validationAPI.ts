import axios from "@/plugins/axios";

// const path = "validate"

export async function Test() {
    const response = await axios.get("ping")
    return response
}