<template>
  <ion-page>
    <!-- <ion-header>
      <ion-toolbar>
        <ion-title>我的e+</ion-title>
      </ion-toolbar>
    </ion-header> -->
    <br />
    <ion-item>
      <ion-thumbnail slot="start">
        <img alt="Silhouette of mountains" src="assets/peizhen3.png" />
      </ion-thumbnail>
      <ion-label>咿呀学语</ion-label>
    </ion-item>
    <ion-content :fullscreen="true">
      <!-- <ion-header collapse="condense">
        <ion-toolbar>
          <ion-title size="large">Tab 4</ion-title>
        </ion-toolbar>
      </ion-header> -->
      <br />
      <ion-item button :detail="true" lines="none" @click="openModal">
        <ion-label>
          <h3>陪诊员入驻</h3>
        </ion-label>
      </ion-item>
      <ion-item button :detail="true" lines="none" @click="openModal">
        <ion-label>
          <h3>服务咨询</h3>
        </ion-label>
      </ion-item>
      <ion-item button :detail="true" lines="none" @click="openModal">
        <ion-label>
          <h3>偏好设置</h3>
        </ion-label>
      </ion-item>
      <ion-item button :detail="true" lines="none" @click="openModal">
        <ion-label>
          <h3>关于我们</h3>
        </ion-label>
      </ion-item>
      <!-- <ExploreContainer name="TODO" /> -->

      <!-- <ion-button expand="block" size="small" @click="openModal">
        陪诊员加入
        <ion-icon slot="end" name="star"></ion-icon>
      </ion-button> -->
      <p>{{ message }}</p>
      <ion-fab slot="fixed" vertical="bottom" horizontal="center">
        <!-- https://github.com/scopewu/qrcode.vue -->
        <qrcode-vue style="margin: 0 0 50px -20px;" value="http://iyou.city/" />
      </ion-fab>
    </ion-content>
  </ion-page>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { IonPage, /*IonHeader, IonToolbar, IonTitle,*/ IonContent, modalController, /*IonButton,*/ IonItem, IonLabel, IonThumbnail, IonFab } from '@ionic/vue';
import { caretForwardOutline } from 'ionicons/icons';
//import ExploreContainer from '@/components/ExploreContainer.vue';
import ModalAttendant from '@/components/ModalAttendant.vue';
import QrcodeVue from 'qrcode.vue'

const message = ref();

async function openModal() {
  const modal = await modalController.create({
    component: ModalAttendant,
  });
  modal.present();

  const { data, role } = await modal.onWillDismiss();
  if (role === 'confirm') {
    message.value = `Hello, ${data}!`;
  }
}
</script>
<style>
ion-thumbnail {
  --size: 60px;
  --border-radius: 30px;
}
</style>