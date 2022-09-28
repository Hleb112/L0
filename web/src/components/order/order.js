import React, {Component} from "react";
import Service from "../../services/service";

export default class Order extends Component {

    service = new Service();

    state = {
        order: null,
        id: null
    };

    constructor(props) {
        super(props);

        this.handleChange = this.handleChange.bind(this);
        this.handleClick = this.handleClick.bind(this);
    };

    onOrderLoaded = (order) => {
        this.setState({order});
    };

    updateOrder(id) {
        this.service
            .getOrderById(id)
            .then(this.onOrderLoaded);
    };

    handleChange(event) {
        this.setState({id: event.target.value});
    };

    handleClick(e) {
        e.preventDefault();
        this.updateOrder(this.state.id);
    };

    renderItems(arr) {
        return Object?.keys(arr).map(key => {
            return (
                <li className="list-group-item"
                    key={key}>
                    {key} : {arr[key]}
                </li>
            );
        });
    };

    // renderItemsOut(arr) {
    //     return
    //     });
    // };

    render() {
        if (this.state.order == null) {
            return (
                <div style={{
                    backgroundColor: 'whitesmoke',
                }}>
                    <form>
                        <label>
                            Введите id заказа:
                            <input type="text" name="order_uuid" onChange={this.handleChange}/>
                        </label>
                        <input type="submit" value="Отправить" onClick={this.handleClick}/>
                    </form>
                </div >
            )
        } else if (this.state.order.message != null) {
            const { message
            } = this.state.order;

            return (
                <div style={{
                    backgroundColor: 'whitesmoke',
                }}>>
                    <form>
                        <label>
                            Введите id заказа:
                            <input type="text" name="order_uuid" onChange={this.handleChange}/>
                        </label>
                        <input type="submit" value="Отправить" onClick={this.handleClick}/>
                        <ul>
                            {"Message"}: {message}
                        </ul>
                    </form>
                </div>
            )
        } else {
            const {
                order: {
                    order_uuid, track_number, entry, delivery, payment, items, locale, internal_signature,
                    customer_id, delivery_service, shard_key, sm_id, date_created, oof_shard
                }
            } = this.state;

            const arr = {
                order_uuid, track_number, entry, locale, internal_signature, customer_id,
                delivery_service, shard_key, sm_id, date_created, oof_shard
            }
            const item = this.renderItems(arr);
            const delivery_out = this.renderItems(delivery)
            const payment_out = this.renderItems(payment)
            //const items_out = this.renderItems(items)
            const myArrMadeFromForEach = [];

            items?.forEach((item, i) => {
                myArrMadeFromForEach.push(<span className="list-group-item">
                    {"Item " + (i + 1) + ": "}
                </span>)
                Object.keys(item).map(key => {
                    console.log(key)
                    console.log(item[key])
                        myArrMadeFromForEach.push(
                            <li  className="list-group-item"
                                key={key}>
                                {key} : {item[key]}
                            </li>
                        )
                    }
                )
            })
            // console.log(items_out)

            return (
                <div style={{
                    backgroundColor: 'whitesmoke',
                }}>>
                    <form>
                        <label>
                            Введите id заказа:
                            <input type="text" name="order_uuid" onChange={this.handleChange}/>
                        </label>
                        <input type="submit" value="Отправить" onClick={this.handleClick}/>
                    </form>
                    <ul className="item-list list-group">
                        "Order" : {item}
                        "Delivery" : {delivery_out}
                        "Payment" : {payment_out}
                        "Items" : <br/>
                        {myArrMadeFromForEach}
                    </ul>
                </div>
            );
        }
    };
};