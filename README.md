This is a microservice application for advisor categories.

Prerequisites
1) Mongodb server running with below database and collection
database: finman-db
collection: category 

2) this requires PMS_CONFIG environment variable set with the complete filepath having below properties
mongodb.host: <host>
mongodb.port: <port> defaults to 27017
mongodb.username: <username>
mongodb.password: <password>

Sample Document
{
    "_id": "1",
    "type": "fee-only",
    "description": "There are numerous Financial Planners and Advisors around the country. They provide professional financial advice and earn for their living. However, how they get paid is where the difference is. This is important because it directly affects the quality advice you get from the planner. Fee-only financial advisors charge a flat fee on all transactions, regardless of the portfolio reading Rs 50 lakh or Rs 5 crore. Unlike fee-based advisors, we can expect more uniformity from the fee-only financial advisors in servicing clients. This is because the fee remains the same for all customers.",
    "advantages": [
        "INDEPENDENT",
        "LICENCED",
        "UNBIASED",
        "PERSONALIZED"
    ],
    "disadvantages": [
        "HIGH FEES",
        "HIGH MINIMUM INVESTMENTS"
    ],
    "faq": [
        {
            "question": "What is fee-only financial planning?",
            "answer": "There are numerous Financial Planners and Advisors around the country. They provide professional financial advice and earn for their living. However, how they get paid is where the difference is. This is important because it directly affects the quality advice you get from the planner. Fee-only financial advisors charge a flat fee on all transactions, regardless of the portfolio reading Rs 50 lakh or Rs 5 crore. Unlike fee-based advisors, we can expect more uniformity from the fee-only financial advisors in servicing clients. This is because the fee remains the same for all customers."
        },
        {
            "question": "Does it make sense to pay fees for advise ?",
            "answer": "Fee-only financial advisors charge a flat fee on all transactions, regardless of the portfolio reading Rs 50 lakh or Rs 5 crore. Unlike fee-based advisors, we can expect more uniformity from the fee-only financial advisors in servicing clients. This is because the fee remains the same for all customers."
        },
        {
            "question": "Putting the investor's interest above the advisor's interest",
            "answer": "Commissions and fees are one aspect, but the most important responsibility of a financial planner is to act as a fiduciary towards their clients. To always put the client’s interest ahead of their own. Unfortunately, where there are commissions involved, biases always come in. Planners are often driven to promote and recommend products that offer them the best commission or are limited by the products of companies that they have a tie up with. They may end up in a position to put their interests ahead of the clients’."
        },
        {
            "question": "Biased / Unbiased, confict free advice ?",
            "answer": "Proper financial advice can be given only where there is no bias and all available options are evaluated for the client equally. That is where fee-only financial planning comes in. In fee-only financial planning, the client pays a flat, transparent and mutually agreed upon fee for the financial planner's services. The planner helps the client build a solid, unbiased, personalized financial plan and recommends appropriate products for the client’s needs. The financial planner does not sell or distribute any of the products that he recommends to the client and the only single source of income he gets is from the fee that the client pays."
        }
    ]
}

Application has below rest endpoints

GET /api/category => this gets all the advisor categories
GET /api/category/{id} => this gets the category for given id
POST /api/category => this creates a new advisor category
PUT /api/category/{id} => this updates the category for given id
DELETE /api/category/{id} => this deletes the category for given id

