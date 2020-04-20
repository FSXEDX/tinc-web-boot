import requests

from dataclasses import dataclass

from typing import Any, Optional, List



@dataclass
class Network:
    name: 'str'
    running: 'bool'
    config: 'Optional[Config]'

    def to_json(self) -> dict:
        return {
            "name": self.name,
            "running": self.running,
            "config": self.config.to_json(),
        }

    @staticmethod
    def from_json(payload: dict) -> 'Network':
        return Network(
                name=payload['name'],
                running=payload['running'],
                config=Config.from_json(payload['config']),
        )


@dataclass
class Config:
    name: 'str'
    port: 'int'
    interface: 'str'
    auto_start: 'bool'
    mode: 'str'
    device_type: 'Optional[str]'
    device: 'Optional[str]'
    connect_to: 'Optional[List[str]]'

    def to_json(self) -> dict:
        return {
            "name": self.name,
            "port": self.port,
            "interface": self.interface,
            "autostart": self.auto_start,
            "mode": self.mode,
            "deviceType": self.device_type,
            "device": self.device,
            "connectTo": self.connect_to,
        }

    @staticmethod
    def from_json(payload: dict) -> 'Config':
        return Config(
                name=payload['name'],
                port=payload['port'],
                interface=payload['interface'],
                auto_start=payload['autostart'],
                mode=payload['mode'],
                device_type=payload['deviceType'],
                device=payload['device'],
                connect_to=payload['connectTo'] or [],
        )


@dataclass
class PeerInfo:
    name: 'str'
    online: 'bool'
    status: 'Optional[Peer]'
    configuration: 'Optional[Node]'

    def to_json(self) -> dict:
        return {
            "name": self.name,
            "online": self.online,
            "status": self.status.to_json(),
            "config": self.configuration.to_json(),
        }

    @staticmethod
    def from_json(payload: dict) -> 'PeerInfo':
        return PeerInfo(
                name=payload['name'],
                online=payload['online'],
                status=Peer.from_json(payload['status']),
                configuration=Node.from_json(payload['config']),
        )


@dataclass
class Peer:
    node: 'str'
    subnet: 'str'
    fetched: 'bool'

    def to_json(self) -> dict:
        return {
            "node": self.node,
            "subnet": self.subnet,
            "fetched": self.fetched,
        }

    @staticmethod
    def from_json(payload: dict) -> 'Peer':
        return Peer(
                node=payload['node'],
                subnet=payload['subnet'],
                fetched=payload['fetched'],
        )


@dataclass
class Node:
    name: 'str'
    subnet: 'str'
    port: 'int'
    address: 'Optional[List[Address]]'
    public_key: 'str'
    version: 'int'

    def to_json(self) -> dict:
        return {
            "name": self.name,
            "subnet": self.subnet,
            "port": self.port,
            "address": [x.to_json() for x in self.address],
            "publicKey": self.public_key,
            "version": self.version,
        }

    @staticmethod
    def from_json(payload: dict) -> 'Node':
        return Node(
                name=payload['name'],
                subnet=payload['subnet'],
                port=payload['port'],
                address=[Address.from_json(x) for x in (payload['address'] or [])],
                public_key=payload['publicKey'],
                version=payload['version'],
        )


@dataclass
class Address:
    host: 'str'
    port: 'Optional[int]'

    def to_json(self) -> dict:
        return {
            "host": self.host,
            "port": self.port,
        }

    @staticmethod
    def from_json(payload: dict) -> 'Address':
        return Address(
                host=payload['host'],
                port=payload['port'],
        )


@dataclass
class Sharing:
    name: 'str'
    nodes: 'Optional[List[Node]]'

    def to_json(self) -> dict:
        return {
            "name": self.name,
            "node": [x.to_json() for x in self.nodes],
        }

    @staticmethod
    def from_json(payload: dict) -> 'Sharing':
        return Sharing(
                name=payload['name'],
                nodes=[Node.from_json(x) for x in (payload['node'] or [])],
        )


@dataclass
class Upgrade:
    subnet: 'Optional[str]'
    port: 'Optional[int]'
    address: 'Optional[List[Address]]'
    device: 'Optional[str]'

    def to_json(self) -> dict:
        return {
            "subnet": self.subnet,
            "port": self.port,
            "address": [x.to_json() for x in self.address],
            "device": self.device,
        }

    @staticmethod
    def from_json(payload: dict) -> 'Upgrade':
        return Upgrade(
                subnet=payload['subnet'],
                port=payload['port'],
                address=[Address.from_json(x) for x in (payload['address'] or [])],
                device=payload['device'],
        )


class TincWebError(RuntimeError):
    def __init__(self, method: str, code: int, message: str, data: Any):
        super().__init__('{}: {}: {} - {}'.format(method, code, message, data))
        self.code = code
        self.message = message
        self.data = data

    @staticmethod
    def from_json(method: str, payload: dict) -> 'TincWebError':
        return TincWebError(
            method=method,
            code=payload['code'],
            message=payload['message'],
            data=payload.get('data')
        )


class TincWebClient:
    """
    Public Tinc-Web API (json-rpc 2.0)
    """

    def __init__(self, base_url: str = 'http://127.0.0.1:8686/api/', session: Optional[requests.Session] = None):
        self.__url = base_url
        self.__id = 1
        self.__session = session or requests

    def __next_id(self):
        self.__id += 1
        return self.__id

    def networks(self) -> List[Network]:
        """
        List of available networks (briefly, without config)
        """
        response = self.__session.post(self.__url, json={
            "jsonrpc": "2.0",
            "method": "TincWeb.Networks",
            "id": self.__next_id(),
            "params": []
        })
        assert response.ok, str(response.status_code) + " " + str(response.reason)
        payload = response.json()
        if 'error' in payload:
            raise TincWebError.from_json('networks', payload['error'])
        return [Network.from_json(x) for x in (payload['result'] or [])]

    def network(self, name: str) -> Network:
        """
        Detailed network info
        """
        response = self.__session.post(self.__url, json={
            "jsonrpc": "2.0",
            "method": "TincWeb.Network",
            "id": self.__next_id(),
            "params": [name, ]
        })
        assert response.ok, str(response.status_code) + " " + str(response.reason)
        payload = response.json()
        if 'error' in payload:
            raise TincWebError.from_json('network', payload['error'])
        return Network.from_json(payload['result'])

    def create(self, name: str) -> Network:
        """
        Create new network if not exists
        """
        response = self.__session.post(self.__url, json={
            "jsonrpc": "2.0",
            "method": "TincWeb.Create",
            "id": self.__next_id(),
            "params": [name, ]
        })
        assert response.ok, str(response.status_code) + " " + str(response.reason)
        payload = response.json()
        if 'error' in payload:
            raise TincWebError.from_json('create', payload['error'])
        return Network.from_json(payload['result'])

    def remove(self, network: str) -> bool:
        """
        Remove network (returns true if network existed)
        """
        response = self.__session.post(self.__url, json={
            "jsonrpc": "2.0",
            "method": "TincWeb.Remove",
            "id": self.__next_id(),
            "params": [network, ]
        })
        assert response.ok, str(response.status_code) + " " + str(response.reason)
        payload = response.json()
        if 'error' in payload:
            raise TincWebError.from_json('remove', payload['error'])
        return payload['result']

    def start(self, network: str) -> Network:
        """
        Start or re-start network
        """
        response = self.__session.post(self.__url, json={
            "jsonrpc": "2.0",
            "method": "TincWeb.Start",
            "id": self.__next_id(),
            "params": [network, ]
        })
        assert response.ok, str(response.status_code) + " " + str(response.reason)
        payload = response.json()
        if 'error' in payload:
            raise TincWebError.from_json('start', payload['error'])
        return Network.from_json(payload['result'])

    def stop(self, network: str) -> Network:
        """
        Stop network
        """
        response = self.__session.post(self.__url, json={
            "jsonrpc": "2.0",
            "method": "TincWeb.Stop",
            "id": self.__next_id(),
            "params": [network, ]
        })
        assert response.ok, str(response.status_code) + " " + str(response.reason)
        payload = response.json()
        if 'error' in payload:
            raise TincWebError.from_json('stop', payload['error'])
        return Network.from_json(payload['result'])

    def peers(self, network: str) -> List[PeerInfo]:
        """
        Peers brief list in network  (briefly, without config)
        """
        response = self.__session.post(self.__url, json={
            "jsonrpc": "2.0",
            "method": "TincWeb.Peers",
            "id": self.__next_id(),
            "params": [network, ]
        })
        assert response.ok, str(response.status_code) + " " + str(response.reason)
        payload = response.json()
        if 'error' in payload:
            raise TincWebError.from_json('peers', payload['error'])
        return [PeerInfo.from_json(x) for x in (payload['result'] or [])]

    def peer(self, network: str, name: str) -> PeerInfo:
        """
        Peer detailed info by in the network
        """
        response = self.__session.post(self.__url, json={
            "jsonrpc": "2.0",
            "method": "TincWeb.Peer",
            "id": self.__next_id(),
            "params": [network, name, ]
        })
        assert response.ok, str(response.status_code) + " " + str(response.reason)
        payload = response.json()
        if 'error' in payload:
            raise TincWebError.from_json('peer', payload['error'])
        return PeerInfo.from_json(payload['result'])

    def import(self, sharing: Sharing) -> Network:
        """
        Import another tinc-web network configuration file.
It means let nodes defined in config join to the network.
Return created (or used) network with full configuration
        """
        response = self.__session.post(self.__url, json={
            "jsonrpc": "2.0",
            "method": "TincWeb.Import",
            "id": self.__next_id(),
            "params": [sharing.to_json(), ]
        })
        assert response.ok, str(response.status_code) + " " + str(response.reason)
        payload = response.json()
        if 'error' in payload:
            raise TincWebError.from_json('import', payload['error'])
        return Network.from_json(payload['result'])

    def share(self, network: str) -> Sharing:
        """
        Share network and generate configuration file.
        """
        response = self.__session.post(self.__url, json={
            "jsonrpc": "2.0",
            "method": "TincWeb.Share",
            "id": self.__next_id(),
            "params": [network, ]
        })
        assert response.ok, str(response.status_code) + " " + str(response.reason)
        payload = response.json()
        if 'error' in payload:
            raise TincWebError.from_json('share', payload['error'])
        return Sharing.from_json(payload['result'])

    def node(self, network: str) -> Node:
        """
        Node definition in network (aka - self node)
        """
        response = self.__session.post(self.__url, json={
            "jsonrpc": "2.0",
            "method": "TincWeb.Node",
            "id": self.__next_id(),
            "params": [network, ]
        })
        assert response.ok, str(response.status_code) + " " + str(response.reason)
        payload = response.json()
        if 'error' in payload:
            raise TincWebError.from_json('node', payload['error'])
        return Node.from_json(payload['result'])

    def upgrade(self, network: str, update: Upgrade) -> Node:
        """
        Upgrade node parameters.
In some cases requires restart
        """
        response = self.__session.post(self.__url, json={
            "jsonrpc": "2.0",
            "method": "TincWeb.Upgrade",
            "id": self.__next_id(),
            "params": [network, update.to_json(), ]
        })
        assert response.ok, str(response.status_code) + " " + str(response.reason)
        payload = response.json()
        if 'error' in payload:
            raise TincWebError.from_json('upgrade', payload['error'])
        return Node.from_json(payload['result'])