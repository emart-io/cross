<template>
  <ion-page>
    <ion-header>
      <ion-toolbar>
        <ion-button size="small" fill="clear" v-on:click="getLocation" slot="start">
          <ion-icon slot="start" :icon="location"></ion-icon>
          {{ position }}
        </ion-button>
        <ion-button slot="end" fill="clear">
          <!-- e+陪诊 -->
        </ion-button>
      </ion-toolbar>
    </ion-header>
    <ion-content :fullscreen="true">
      <swiper :modules="modules" :slides-per-view="1" :space-between="50" :autoplay="autoplay" @swiper="onSwiper"
        @slideChange="onSlideChange" :pagination="{ clickable: true }">
        <swiper-slide>
          <!-- <ion-header collapse="condense">
                                                          <ion-toolbar>
                                                            <ion-title size="large">Tab 1</ion-title>
                                                          </ion-toolbar>
                                                        </ion-header>
                                                        <ExploreContainer name="Tab 1 page" /> -->
          <img alt="Silhouette of mountains" src="assets/0211.jpeg" />
        </swiper-slide>
        <swiper-slide>
          <!-- <ion-title>暖心陪诊</ion-title> -->
          <img alt="Silhouette of mountains" src="assets/bujk.jpeg" />
        </swiper-slide>
        <swiper-slide>
          <img alt="Silhouette of mountains" src="assets/peizhen5.png" />
        </swiper-slide>
      </swiper>
      <div id="container" style="display:none"></div>
      <ion-grid>
        <ion-row>
          <ion-col>
            <ion-card>
              <img alt="Silhouette of mountains" src="assets/pz-ser-banner.jpg" />
              <ion-card-header>
                <ion-card-title>就医陪诊</ion-card-title>
                <ion-card-subtitle>暖心陪诊，安心</ion-card-subtitle>
              </ion-card-header>
            </ion-card>
          </ion-col>
          <ion-col>
            <ion-card>
              <img alt="Silhouette of mountains" src="assets/pz-ser-banner.jpg" />
              <ion-card-header>
                <ion-card-title>取送结果</ion-card-title>
                <ion-card-subtitle>安全方便，快捷</ion-card-subtitle>
              </ion-card-header>
            </ion-card>
          </ion-col>
        </ion-row>
      </ion-grid>
      <ion-list>
        <ion-item v-for="item in hospitals" :key="item['id']" lines="none">
          <ion-thumbnail slot="start">
            <img alt="Silhouette of mountains" :src="item['photos'][0]['url']" style="border-radius: 10px" />
          </ion-thumbnail>
          <ion-label>
            <h3>{{ item['name'] }}</h3>
            <p>地址:{{ item['address'] }}</p>
            <p>电话:{{ item['tel'] }}</p>
          </ion-label>
        </ion-item>
      </ion-list>
    </ion-content>
</ion-page>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { IonPage, IonHeader, IonToolbar, IonButton, IonIcon, IonList, IonLabel, IonContent, IonCard, /*IonCardContent,*/ IonCardHeader, IonCardSubtitle, IonCardTitle, IonCol, IonGrid, IonRow, IonItem, IonThumbnail } from '@ionic/vue';
//import ExploreContainer from '@/components/ExploreContainer.vue';
import { location } from 'ionicons/icons';
import AMapLoader from '@amap/amap-jsapi-loader';
import { Autoplay, Pagination } from 'swiper';
import { Swiper, SwiperSlide } from 'swiper/vue';

import 'swiper/css';
import 'swiper/css/autoplay';
import 'swiper/css/pagination';

export default defineComponent({
  name: 'Tab1Page',
  ionViewWillEnter() {
    console.log('Home page will enter');
    this.getLocation();
  },
  components: { /*ExploreContainer,*/ IonHeader, IonToolbar, IonButton, IonIcon, IonList, IonLabel, IonContent, IonPage, Swiper, SwiperSlide, IonCard, /*IonCardContent,*/ IonCardHeader, IonCardSubtitle, IonCardTitle, IonCol, IonGrid, IonRow, IonItem, IonThumbnail },
  data() {
    return {
      position: "湖北荆门",
      autoplay: { delay: 3000 },
      hospitals: [],
    }
  },
  methods: {
    getLocation() {
      AMapLoader.load({
        key: '60d396703bef1a6a93d2eca45a70e764',
        version: '1.4.15',
        plugins: ['AMap.Geolocation', 'AMap.PlaceSearch'],
      }).then(AMap => {
        const geolocation = new AMap.Geolocation({
          enableHighAccuracy: true,
          radius: 10000,
        })

        geolocation.getCurrentPosition((status: any, result: any) => {
          if (status == 'complete') {
            console.log("111", result)
            if (result.addressComponent) {
              this.position = result.addressComponent.province;
            }
          } else {
            console.log("222", result);
          }
        })

        var placeSearch = new AMap.PlaceSearch({
          pageSize: 100, // 单页显示结果条数
          // city 指定搜索所在城市，支持传入格式有：城市名、citycode和adcode
          city: '010',
          panel: "container",
        });

        placeSearch.search('三甲医院', (status: any, result: any) => {
          // 查询成功时，result即对应匹配的POI信息
          console.log(JSON.stringify(result));
          this.$data.hospitals = result.poiList.pois;
        });
      }).catch((e) => {
        console.log(e)
      })
    },
  },
  setup() {
    const onSwiper = (swiper: any) => {
      console.log(swiper);
    };
    const onSlideChange = () => {
      console.log('slide change');
    };
    // const map = shallowRef(null);
    return {
      location,
      onSwiper,
      onSlideChange,
      modules: [Autoplay, Pagination],
      // map
    }
  }
});
</script>

<style>
.swiper {
  width: 100%;
  height: 33%;
}

.swiper .swiper-slide {
  height: auto;
}

.swiper-slide img {
  display: block;
  width: 100%;
  height: 100%;
  object-fit: cover;
  /* border-radius: 15px; */
}
</style>