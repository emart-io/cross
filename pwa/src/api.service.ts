import { AttendantsPromiseClient } from './sdk/attendant_grpc_web_pb';

export class ApiService {
    // opts = { 'streamInterceptors': [new StreamInterceptor()] };;
    apiUrl = 'https://' + window.location.host

    attendantClient = new AttendantsPromiseClient(this.apiUrl)
    // commodityClient = new CommoditiesPromiseClient(environment.apiUrl, null, this.opts);
    // userClient = new UsersPromiseClient(environment.apiUrl);
    // couponClient = new CouponsPromiseClient(environment.apiUrl);
    // orderClient = new OrdersPromiseClient(environment.apiUrl);
    // addressClient = new AddressesPromiseClient(environment.apiUrl);
    // messageClient = new MessagesPromiseClient(environment.apiUrl);
    // accountClient = new AccountsPromiseClient(environment.apiUrl);
    // commentClient = new CommentsPromiseClient(environment.apiUrl);
    // memoClient = new MemosPromiseClient(environment.apiUrl);

    //constructor() { }

}

export const apiService = new ApiService();