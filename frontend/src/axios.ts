import axios from "axios";

export default axios.create({
	baseURL: 'https://flood.mcax.cn/api/',
	headers: {
		'Content-Type': 'application/json',
	},
})
