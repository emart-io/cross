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
    <!-- <ion-content class="ion-padding"> -->
    <ion-content>
        <ion-item>
            <ion-label>就诊医院</ion-label>
            <!-- <ion-input v-model="order.location" placeholder="请输入医院名称" required></ion-input> -->
            <ion-select mode="ios" aria-label="就诊医院" interface="action-sheet" placeholder="请选择就诊医院">
                <ion-select-option value="apples">Apples</ion-select-option>
                <ion-select-option value="oranges">Oranges</ion-select-option>
                <ion-select-option value="bananas">Bananas</ion-select-option>
            </ion-select>
        </ion-item>
        <ion-item>
            <ion-label>就诊科室</ion-label>
            <ion-input slot="end" v-model="order.name" placeholder="请输入就诊科室" required></ion-input>
        </ion-item>
        <ion-item>
            <ion-label>就诊时间</ion-label>
            <ion-input slot="end" v-model="order.name" placeholder="请选择就诊时间" required></ion-input>
        </ion-item>
        <ion-item>
            <ion-label>就诊人</ion-label>
            <ion-input slot="end" v-model="order.name" placeholder="请选择就诊人" required></ion-input>
        </ion-item>
        <ion-item>
            <ion-label>备注</ion-label>
            <ion-textarea slot="end" v-model="order.description" :rows="lines"></ion-textarea>
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
    IonSelect,
    IonSelectOption,
    modalController,
} from '@ionic/vue';
import { defineComponent } from 'vue';
import { apiService } from '../api.service';
import { Order } from '../sdk/order_pb';

export default defineComponent({
    name: 'ModalOrder',
    components: { IonContent, IonHeader, IonTitle, IonToolbar, IonButtons, IonButton, IonItem, IonLabel, IonInput, IonTextarea, IonSelect, IonSelectOption },
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