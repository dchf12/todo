import axios from 'axios';

const API_HOST = 'http://localhost:8080';

export default axios.create({
  baseURL: API_HOST,
});
