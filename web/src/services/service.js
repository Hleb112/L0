export default class Service {

    _apiBase = 'http://localhost:8080/api/';

    getResource = async (url) => {
        const res = await fetch(`${this._apiBase}${url}`);

        if (res.status !== 200 || !res.ok) {
            return res.json();
        }

        return await res.json();
    }

    getOrderById = async (id) => {
        const order = await this.getResource(`v1/order/${id}`)
        if (order.message != null) {
            return this._transformBadRequest(order);
        }

        return this._transformOrder(order);
    };

    _transformBadRequest= (res) => {
        return {
            message: res.message
        };
    };

    _transformOrder = (order) => {
        return {
            order_uuid: order.order_uuid,
            track_number: order.track_number,
            entry: order.entry,
            locale: order.locale,
            internal_signature: order.internal_signature,
            customer_id: order.customer_id,
            delivery_service: order.delivery_service,
            shard_key: order.shardkey,
            sm_id: order.sm_id,
            date_created: order.date_created,
            oof_shard: order.oof_shard,
            delivery: order.delivery,
            payment: order.payment,
            items: order.items,
        };
    };
};