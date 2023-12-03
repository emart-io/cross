<template>
    <ion-header>
        <ion-toolbar>
            <ion-title>陪诊服务下单</ion-title>
        </ion-toolbar>
    </ion-header>
    <ion-content>
        <ion-card>
            <ion-card-header>
                <ion-card-title></ion-card-title>
                <ion-card-subtitle>陪诊方式</ion-card-subtitle>
            </ion-card-header>
            <ion-card-content>
                <ion-item>
                    <ion-select interface="popover" placeholder="请选择陪诊方式" v-model="order.type" required>
                        <ion-select-option value="全天">全天陪诊</ion-select-option>
                        <ion-select-option value="半天">半天陪诊</ion-select-option>
                        <ion-select-option value="取送结果">取送结果</ion-select-option>
                    </ion-select>
                </ion-item>
            </ion-card-content>
            <ion-card-header>
                <ion-card-title></ion-card-title>
                <ion-card-subtitle>就诊信息</ion-card-subtitle>
            </ion-card-header>
            <ion-card-content>
                <ion-item>
                    <ion-select label="就诊医院" interface="action-sheet" placeholder="请选择就诊医院" v-model="order.hospital"
                        required>
                        <ion-select-option v-for="item in hospitals" :value="item['name']" :key="item['id']">{{ item['name']
                        }}</ion-select-option>
                        <!-- <ion-select-option value="oranges">Oranges</ion-select-option> -->
                    </ion-select>
                </ion-item>
                <ion-item>
                    <ion-input label="就诊科室" v-model="order.department" placeholder="请输入就诊科室" required></ion-input>
                </ion-item>
                <ion-item>
                    <ion-label>就诊时间</ion-label>
                    <ion-datetime-button datetime="datetime" :show-default-buttons="true"></ion-datetime-button>
                    <ion-modal :keep-contents-mounted="true">
                        <ion-datetime id="datetime"></ion-datetime>
                    </ion-modal>
                </ion-item>
                <ion-item>
                    <ion-input label="就诊人" v-model="order.patient" placeholder="请输入就诊人" required></ion-input>
                </ion-item>
                <ion-item>
                    <ion-textarea label="备注" v-model="order.notes" :rows="lines" placeholder="添加就诊注意事项"></ion-textarea>
                </ion-item>
            </ion-card-content>
            <ion-toolbar>
                <ion-buttons slot="start">
                    <ion-button color="medium" @click="cancel">取消</ion-button>
                </ion-buttons>
                <ion-buttons slot="end">
                    <ion-button @click="confirm">确定</ion-button>
                </ion-buttons>
            </ion-toolbar>
        </ion-card>
    </ion-content>
</template>
  
<script setup lang="ts">
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
    IonDatetime, IonDatetimeButton, IonModal,
    modalController,
    IonCard, IonCardContent, IonCardHeader, IonCardSubtitle, IonCardTitle
} from '@ionic/vue';
import { apiService } from '../api.service';
import { Order } from '../sdk/order_pb';

var lines = 2;
var order = new Order().toObject();
var hospitals = apiService.hospitals;

function cancel() {
    return modalController.dismiss(null, 'cancel');
}
function confirm() {
    console.log(order);
    if (order.hospital == '' || order.department == '') {
        alert('请选择医院与科室')
        return
    }
    var att = new Order().setHospital(order.hospital).setDepartment(order.department);
    att.setNotes(order.notes);
    apiService.orderClient.add(att).catch(err => {
        console.log(JSON.stringify(err));
    });
    return modalController.dismiss(order.hospital, 'confirm');
}
</script>