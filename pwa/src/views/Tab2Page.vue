<template>
  <ion-page>
    <ion-header>
      <ion-toolbar style="--background: linear-gradient(to right,yellow,white);">
        <ion-title>暖心陪诊</ion-title>
      </ion-toolbar>
    </ion-header>
    <ion-content :fullscreen="true">
      <ion-header collapse="condense">
        <ion-toolbar>
          <ion-title size="large">Tab 2</ion-title>
        </ion-toolbar>
      </ion-header>
      {{ user.name }}
      <input v-model="user.name" />
      <ExploreContainer name="Tab 2 page " />
    </ion-content>
  </ion-page>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { IonPage, IonHeader, IonToolbar, IonTitle, IonContent } from '@ionic/vue';
import ExploreContainer from '@/components/ExploreContainer.vue';
import { User } from '../sdk/user_pb';

export default defineComponent({
  name: 'Tab2Page',
  components: { ExploreContainer, IonHeader, IonToolbar, IonTitle, IonContent, IonPage },
  data() {
    var user = new User().toObject();
    user.name = 'test';

    //let u2=new User();
    new User().setName(user.name).setIcon("").setId("d");
    return {
      user
    }
  },
});
</script>
