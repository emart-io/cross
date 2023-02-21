<template>
  <ion-page>
    <ion-header>
      <ion-toolbar>
        <ion-title>服务订单</ion-title>
      </ion-toolbar>
    </ion-header>
    <ion-content :fullscreen="true">
      <!-- <ion-header collapse="condense">
        <ion-toolbar>
          <ion-title size="large">Tab 3</ion-title>
        </ion-toolbar>
      </ion-header>
      
      <ExploreContainer name="Tab 3 page" /> -->

      <ion-segment :scrollable="true" value="default">
        <ion-segment-button value="default">
          <ion-label>全部</ion-label>
        </ion-segment-button>
        <ion-segment-button value="default1">
          <ion-label>待支付</ion-label>
        </ion-segment-button>
        <ion-segment-button value="default2">
          <ion-label>待接单</ion-label>
        </ion-segment-button>
        <ion-segment-button value="default3">
          <ion-label>进行中</ion-label>
        </ion-segment-button>
        <ion-segment-button value="default4">
          <ion-label>已完成</ion-label>
        </ion-segment-button>
      </ion-segment>
    </ion-content>
  </ion-page>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { IonPage, IonHeader, IonToolbar, IonTitle, IonContent, IonLabel, IonSegment, IonSegmentButton } from '@ionic/vue';
//import ExploreContainer from '@/components/ExploreContainer.vue';

export default defineComponent({
  name: 'Tab3Page',
  components: { /*ExploreContainer,*/ IonHeader, IonToolbar, IonTitle, IonContent, IonPage, IonLabel, IonSegment, IonSegmentButton }
});
</script>
