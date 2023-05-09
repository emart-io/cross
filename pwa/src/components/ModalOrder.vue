<template>
    <ion-header>
        <ion-toolbar>
            <ion-buttons slot="start">
                <ion-button color="medium" @click="cancel">取消</ion-button>
            </ion-buttons>
            <ion-title>陪诊服务下单</ion-title>
            <ion-buttons slot="end">
                <ion-button @click="confirm">确定</ion-button>
            </ion-buttons>
        </ion-toolbar>
    </ion-header>
    <ion-content class="ion-padding">
        <ion-item>
            <ion-label>名称：</ion-label>
            <ion-input v-model="order.name" placeholder="请输入订单名" required></ion-input>
        </ion-item>
        <ion-item>
            <ion-label>地点：</ion-label>
            <ion-input v-model="order.location" placeholder="请输入医院名称" required></ion-input>
        </ion-item>
        <ion-item>
            <ion-label>附加信息：</ion-label>
            <ion-textarea v-model="order.description" placeholder="可选" :rows="lines"></ion-textarea>
        </ion-item>
    </ion-content>
</template>
  
<script lang="ts">
import {
    IonContent,
    IonHeader,
    IonTitle,
    IonToolbar,
    IonButtons,
    IonButton,
    IonItem,
    IonLabel,
    IonInput,
    IonTextarea,
    modalController,
} from '@ionic/vue';
import { defineComponent } from 'vue';
import { apiService } from '../api.service';
import { Order } from '../sdk/order_pb';

export default defineComponent({
    name: 'ModalOrder',
    components: { IonContent, IonHeader, IonTitle, IonToolbar, IonButtons, IonButton, IonItem, IonLabel, IonInput, IonTextarea },
    data() {
        var lines = 2;
        var order = new Order().toObject();
        return {
            order, lines,
        }
    },
    methods: {
        cancel() {
            return modalController.dismiss(null, 'cancel');
        },
        confirm() {
            if (this.order.name == '') {
                alert('请输入名称')
                return
            }
            var att = new Order().setName(this.order.name).setLocation(this.order.location);
            att.setDescription(this.order.description);
            apiService.orderClient.add(att).catch(err => {
                console.log(JSON.stringify(err));
            });
            return modalController.dismiss(this.order.name, 'confirm');
        },
    },
});
</script>