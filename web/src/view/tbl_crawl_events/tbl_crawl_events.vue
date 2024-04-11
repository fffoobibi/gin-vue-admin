<template>
  <el-row justify="center">
    <el-card class="report-card">
      <div class="data-cont">
        <el-scrollbar>
          <ul class="data-news">
            <li>
              <div class="news-title" title="印度服务器进行全面清洗">印度服务器进行全面清洗</div>
              <p>
                3月17日印度新服务器新增资源数突破1器新增资源数突破100W同比上升20%3月17日印度新服务器新增资源数突破1器新增资源数突破100W同比上升20%...3月17日印度新服务器新增资源数突破1器新增资源数突破100W同比上升20%...3月17日印度新服务器新增资源数突破1器新增资源数突破100W同比上升20%...
              </p>
              <time>2024.3.18</time>
            </li>
            <li>
              <div class="news-title" title="3月19日新增一台印度服务器印度新服务器新增资源数突破1器新增资源数突破100W同比上升">
                3月19日新增一台印度服务器印度新服务器新增资源数突破1器新增资源数突破100W同比上升</div>
              <p>3月17日印度新服务器新增资源数突破1器新增资源数突破100W同比上升20%3月17日印度新服务器新增资源数突破1器新增资源数突破100W同比上升20%...</p>
              <time>2024.3.18</time>
            </li>
            <li>
              <div class="news-title" title="印度服务器进行全面清洗">印度服务器进行全面清洗</div>
              <p>
                3月17日印度新服务器新增资源数突破1器新增资源数突破100W同比上升20%3月17日印度新服务器新增资源数突破1器新增资源数突破100W同比上升20%...3月17日印度新服务器新增资源数突破1器新增资源数突破100W同比上升20%...3月17日印度新服务器新增资源数突破1器新增资源数突破100W同比上升20%...
              </p>
              <time>2024.3.18</time>
            </li>
            <li>
              <div class="news-title" title="3月19日新增一台印度服务器印度新服务器新增资源数突破1器新增资源数突破100W同比上升">
                3月19日新增一台印度服务器印度新服务器新增资源数突破1器新增资源数突破100W同比上升</div>
              <p>3月17日印度新服务器新增资源数突破1器新增资源数突破100W同比上升20%3月17日印度新服务器新增资源数突破1器新增资源数突破100W同比上升20%...</p>
              <time>2024.3.18</time>
            </li>
            <li>
              <div class="news-title" title="印度服务器进行全面清洗">印度服务器进行全面清洗</div>
              <p>
                3月17日印度新服务器新增资源数突破1器新增资源数突破100W同比上升20%3月17日印度新服务器新增资源数突破1器新增资源数突破100W同比上升20%...3月17日印度新服务器新增资源数突破1器新增资源数突破100W同比上升20%...3月17日印度新服务器新增资源数突破1器新增资源数突破100W同比上升20%...
              </p>
              <time>2024.3.18</time>
            </li>
            <li>
              <div class="news-title" title="3月19日新增一台印度服务器印度新服务器新增资源数突破1器新增资源数突破100W同比上升">
                3月19日新增一台印度服务器印度新服务器新增资源数突破1器新增资源数突破100W同比上升</div>
              <p>3月17日印度新服务器新增资源数突破1器新增资源数突破100W同比上升20%3月17日印度新服务器新增资源数突破1器新增资源数突破100W同比上升20%...</p>
              <time>2024.3.18</time>
            </li>
            <li>
              <div class="news-title" title="印度服务器进行全面清洗">印度服务器进行全面清洗</div>
              <p>
                3月17日印度新服务器新增资源数突破1器新增资源数突破100W同比上升20%3月17日印度新服务器新增资源数突破1器新增资源数突破100W同比上升20%...3月17日印度新服务器新增资源数突破1器新增资源数突破100W同比上升20%...3月17日印度新服务器新增资源数突破1器新增资源数突破100W同比上升20%...
              </p>
              <time>2024.3.18</time>
            </li>
            <li>
              <div class="news-title" title="3月19日新增一台印度服务器印度新服务器新增资源数突破1器新增资源数突破100W同比上升">
                3月19日新增一台印度服务器印度新服务器新增资源数突破1器新增资源数突破100W同比上升</div>
              <p>3月17日印度新服务器新增资源数突破1器新增资源数突破100W同比上升20%3月17日印度新服务器新增资源数突破1器新增资源数突破100W同比上升20%...</p>
              <time>2024.3.18</time>
            </li>
          </ul>
        </el-scrollbar>

      </div>
      <el-pagination background layout="prev, pager, next" :total="1000" small />
    </el-card>

  </el-row>


</template>

<script setup>


defineOptions({
  name: 'TblCrawlEvents'
})

const getCrawlEvents = async () => {

}



</script>

<style scoped>
.report-card {
  width: 1290px;
  max-height: calc(100vh - 124px);
  /* margin-top: 100px; */
  /* min-height: 800px; */
}

.data-cont {
  text-align: center;
  font-weight: bold;
  font-size: 18px;
  padding-bottom: 31px;
  padding-top: 12px;
  height: calc(100vh - 264px);
}

.data-cont {
  color: #666;
  text-align: center;
  padding-top: 16px;
  font-size: 14px;
  font-weight: normal
}

.data-news {
  padding-left: 20px;
}

.data-news p {
  color: #666;
  padding-top: 6px;
  font-size: 14px;
  margin: 0px;
  text-align: left
}

.data-news li {
  margin-bottom: 28px
}

.data-news li::marker {
  color: #9f9f9f
}

.data-news time {
  display: flex;
  justify-content: left;
  font-size: 12px;
  color: #666;
  position: relative;
  top: 4px;
}

/* .data-news  */
.news-title {
  font-weight: bold;
  text-align: left;
}
</style>
