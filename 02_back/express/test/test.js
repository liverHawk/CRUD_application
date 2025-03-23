const request = require("supertest");
const { expect } = require("chai");

const app = require("../index");

describe("/", () => {
    var agent = request.agent("http://localhost:3000");

    it("health ok", () => {
        agent
            .get("/")
            .set("Accept", "application/json")
            .expect((res) => {
                res.status.should.equal(200);
                expect(res.body.body).to.equal("Hello, World!");
            })
            .end((err, res) => {
                if (err) {
                    throw err;
                }
            });
    });
});
