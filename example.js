import axios from "axios";

async function getArticles() {
  const res = await axios.get("https://discardapi.dpdns.org/api/aljazeera", {
    params: { apikey: "YOUR_API_KEY" }
  });
  console.log(res.data);
}

async function getArticleDetails() {
  const res = await axios.get("https://discardapi.dpdns.org/api/aljazeera/article", {
    params: {
      apikey: "YOUR_API_KEY",
      url: "https://www.aljazeera.com/news/2025/09/20/article-id"
    }
  });
  console.log(res.data);
}

getArticles();
getArticleDetails();
